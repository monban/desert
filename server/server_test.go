package server

import (
	"context"
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/matryer/is"
	"github.com/monban/desert"
	"github.com/nats-io/nats-server/v2/test"
	"github.com/nats-io/nats.go"
)

func RunServer(fn func(*nats.Conn)) {
	s := test.RunRandClientPortServer()
	nc, _ := nats.Connect(s.ClientURL())
	fn(nc)
	nc.Close()
	s.Shutdown()
}

func TestCreateDeckBroadcaster(t *testing.T) {
	// Setup mocks
	is := is.New(t)
	mp := &MockPublisher{}
	daw := &MockDaWatcher{}
	da := desert.DeckEvent{Action: "draw", Card: &desert.Card{}}
	daJson, _ := json.Marshal(da)

	// Create a deck broadcaster
	createDeckBroadcaster(mp, desert.GameId(0), daw, "foo")
	// Ensure something was passed to the Watch function
	is.True(daw.calls.Watch.receives.fn != nil)

	// A function was passed to Watch, call it now
	daw.calls.Watch.receives.fn(da)

	// Check subject
	is.Equal("GAME.0.DECKS.foo", mp.calls.Publish.receives.subj)

	// Check data
	is.Equal(string(daJson), string(mp.calls.Publish.receives.data))
}

func TestSomething(t *testing.T) {
	is := is.New(t)
	gameAction := desert.GameAction{Action: "stormdraw"}
	RunServer(func(nc *nats.Conn) {
		ec, _ := nats.NewEncodedConn(nc, "json")
		nc.Subscribe(">", func(msg *nats.Msg) { t.Log("nats subj:", msg.Subject, "nats msg: ", string(msg.Data)) })
		wait := make(chan struct{})
		ctx, done := context.WithTimeout(context.Background(), time.Second)
		defer done()
		ec.Subscribe("GAME.0.DECKS.>", func(evt desert.DeckEvent) {
			t.Logf("%+v", evt)
			wait <- struct{}{}
		})

		// Start the server
		Run(ctx, nc)

		// Create the game
		reqData, _ := json.Marshal(desert.NewGameData{Name: "foo"})
		_, err := nc.RequestWithContext(ctx, "GAMES.NEW", reqData)
		if err != nil {
			log.Fatal(err)
		}

		// Draw a card
		actionData, _ := json.Marshal(gameAction)
		nc.Publish("GAME.0.ACTION", actionData)

		// See if messages are published on the deck subject
		select {
		case <-wait:
			t.Log("the function was called")
		case <-ctx.Done():
			t.Log("function was not called")
			is.Fail()
		}
	})
}
