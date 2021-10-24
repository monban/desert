package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/monban/desert"
	router "github.com/monban/nats.router"
	"github.com/nats-io/nats.go"
)

func Run(ctx context.Context) error {
	gm := &desert.GameManager{}
	opts := nats.Options{
		Name: "desert",
	}
	nc, err := opts.Connect()
	if err != nil {
		return err
	}
	r := router.Router{Routes: []router.Route{
		{"GAMES.NEW", createNewGameHandler(gm, nc)},
		{"GAME.*.ACTION", createGameActionHandler(gm)},
	}}
	r.ListenAndHandle(ctx, nc)
	<-ctx.Done()
	log.Println("Draining connection")
	nc.Drain()
	return nil
}

func createGameActionHandler(gm *desert.GameManager) router.HandlerFunc {
	return func(ctx context.Context, msg *nats.Msg) {
		gid, _ := strconv.Atoi(strings.Split(msg.Subject, ".")[1])
		game := gm.FindGame(desert.GameId(gid))
		if game == nil {
			log.Printf("Unable to find game with id %d!", gid)
			return
		}

		action := &desert.GameAction{}
		err := json.Unmarshal(msg.Data, action)
		if err != nil {
			log.Printf("Error decoding json %s to action", msg.Data)
		}
		game.HandleAction(*action)
	}
}

func createNewGameHandler(gm *desert.GameManager, nc *nats.Conn) router.HandlerFunc {
	return func(ctx context.Context, msg *nats.Msg) {
		id, _ := gm.NewGame(string(msg.Data))
		g := gm.FindGame(id)
		createDeckBroadcaster(nc, id, &g.StormDeck, "STORM")
		createDeckBroadcaster(nc, id, &g.StormDiscard, "STORM_DISCARD")
		createDeckBroadcaster(nc, id, &g.GearDeck, "GEAR")
		createDeckBroadcaster(nc, id, &g.GearDiscard, "GEAR_DISCARD")
		msg.Respond([]byte(fmt.Sprintf("%+v", id)))
	}
}

func createCardDrawHandler(gm *desert.GameManager) router.HandlerFunc {
	return func(ctx context.Context, msg *nats.Msg) {
		log.Println("Drawing a card")
		sub := strings.Split(msg.Subject, ".")
		id, _ := strconv.ParseUint(sub[1], 10, 64)
		gid := desert.GameId(id)
		deck := sub[3]
		log.Printf("Deck: %v", deck)
		g := gm.FindGame(gid)
		var card desert.Card
		switch deck {
		case "storm":
			card = g.DrawStormCard()
		case "gear":
			card = g.DrawGearCard()
		}
		msg.Respond([]byte(fmt.Sprintf("%+v", card)))
	}
}

func createDeckBroadcaster(nc *nats.Conn, gid desert.GameId, d *desert.Deck, deckName string) {
	fn := func(a desert.DeckAction) {
		subj := fmt.Sprintf("GAME.%d.DECKS.%v", gid, deckName)
		msg := fmt.Sprintf("%+v", a)
		nc.Publish(subj, []byte(msg))
	}
	d.Watch(fn)
}
