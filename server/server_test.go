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
	da := desert.DeckAction{Action: "draw", Card: &desert.Card{}}
	daJson, _ := json.Marshal(da)

	// Create a deck broadcaster
	createDeckBroadcaster(mp, desert.GameId(0), daw, "foo")
	// Ensure something was passed to the Watch function
	is.True(daw.calls.Watch.receives.fn != nil)

	// Call whatever was passed to Watch
	daw.calls.Watch.receives.fn(da)

	// Check subject
	is.Equal("GAME.0.DECKS.foo", mp.calls.Publish.receives.subj)

	// Check data
	is.Equal(string(daJson), string(mp.calls.Publish.receives.data))
}

func TestSomething(t *testing.T) {
	RunServer(func(nc *nats.Conn) {
		nc.Subscribe(">", func(msg *nats.Msg) { t.Log("nats subj:", msg.Subject, "nats msg: ", string(msg.Data)) })
		wait := make(chan struct{})
		type action struct {
			Action string
		}
		//is := is.New(t)
		ctx, done := context.WithTimeout(context.Background(), time.Second*2)
		defer done()
		nc.Subscribe("GAMES.0.DECKS.>", func(msg *nats.Msg) {
			t.Logf("%+v", msg.Data)
			wait <- struct{}{}
		})

		// Start the server
		Run(ctx, nc)

		// Create the game
		reqData, _ := json.Marshal(desert.NewGameData{Name: "foo"})
		res, err := nc.RequestWithContext(ctx, "GAMES.NEW", reqData)
		if err != nil {
			log.Fatal(err)
		}
		t.Log("response: ", string(res.Data))

		// Draw a card
		a := action{"drawstorm"}
		actionData, _ := json.Marshal(a)
		nc.Publish("GAME.0.ACTION", actionData)

		// See if messages are published on the deck subject
		select {
		case <-wait:
			t.Log("the function was called")
		case <-ctx.Done():
			t.Fatal("function was not called")
		}
	})
}
