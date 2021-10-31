package server

import (
	"context"
	"testing"
	"time"

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

func TestSomething(t *testing.T) {
	RunServer(func(nc *nats.Conn) {
		ec, _ := nats.NewEncodedConn(nc, "json")
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
		go Run(ctx, nc)

		var res interface{}
		a := action{"drawstorm"}
		ec.RequestWithContext(ctx, "GAMES.NEW", desert.NewGameData{Name: "foo"}, res)
		ec.Publish("GAME.0.ACTION", a)

		ngd := desert.NewGameData{Name: "Foo"}
		var response []byte
		var foo interface{}
		ec.Request("GAMES.NEW", ngd, foo, time.Second*2)
		nc.Request("GAMES.LIST", response, time.Second*2)
		t.Log(string(response))
		select {
		case <-wait:
			t.Log("the function was called")
		case <-ctx.Done():
			t.Fatal("function was not called")
		}
	})
}
