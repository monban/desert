package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/monban/desert"
	"github.com/nats-io/nats.go"
)

func Run(ctx context.Context, nc *nats.Conn) {
	gm := &desert.GameManager{}
	nc.Subscribe("GAMES.NEW", createNewGameHandler(nc, gm))
	nc.Subscribe("GAMES.LIST", createGameListHandler(nc, gm))
	nc.Subscribe("GAME.*.ACTION", createGameActionHandler(nc, gm))
}

func createGameActionHandler(nc *nats.Conn, gm *desert.GameManager) nats.MsgHandler {
	return func(msg *nats.Msg) {
		gid, _ := strconv.Atoi(strings.Split(msg.Subject, ".")[1])
		game := gm.FindGame(desert.GameId(gid))
		if game == nil {
			log.Printf("Unable to find game with id %d!", gid)
			return
		}

		action := &desert.GameAction{}
		err := json.Unmarshal(msg.Data, action)
		log.Printf("Game %d action: %+v\n", game.Id, action)
		if err != nil {
			log.Printf("Error decoding json %s to action", msg.Data)
		}
		game.HandleAction(*action)
	}
}

func createNewGameHandler(nc Publisher, gm *desert.GameManager) nats.MsgHandler {
	return func(msg *nats.Msg) {
		var ngd desert.NewGameData
		if err := json.Unmarshal(msg.Data, &ngd); err != nil {
			log.Printf("Unable to unmarshal %s to new game data", string(msg.Data))
			return
		}
		log.Printf("Creating a new game: %+v", ngd)
		g, _ := gm.NewGame(ngd)
		createDeckBroadcaster(nc, g.Id, g.StormDeck, "STORM")
		createDeckBroadcaster(nc, g.Id, g.StormDiscard, "STORM_DISCARD")
		createDeckBroadcaster(nc, g.Id, g.GearDeck, "GEAR")
		createDeckBroadcaster(nc, g.Id, g.GearDiscard, "GEAR_DISCARD")
		msg.Respond([]byte(fmt.Sprintf("%+v", g.Id)))
	}
}

type Publisher interface {
	Publish(string, []byte) error
}

type DaWatcher interface {
	Watch(func(desert.DeckAction))
}

func createDeckBroadcaster(nc Publisher, gid desert.GameId, d DaWatcher, deckName string) {
	fn := func(a desert.DeckAction) {
		subj := fmt.Sprintf("GAME.%d.DECKS.%v", gid, deckName)
		log.Printf("Sending events for game %d deck %s to subject %s", gid, deckName, subj)
		msg := fmt.Sprintf("%+v", a)
		nc.Publish(subj, []byte(msg))
	}
	d.Watch(fn)
}
func createGameListHandler(nc *nats.Conn, gm *desert.GameManager) nats.MsgHandler {
	return func(msg *nats.Msg) {
		data, err := json.Marshal(gm.ListGames())
		if err != nil {
			log.Println(err)
			return
		}
		msg.Respond(data)
	}
}
