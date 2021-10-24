package server

import (
	"context"
	"testing"
	"time"

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
		wait := make(chan struct{})
		type action struct {
			Action string
		}
		ec, _ := nats.NewEncodedConn(nc, "json")
		//is := is.New(t)
		ctx, done := context.WithTimeout(context.Background(), time.Second*2)
		defer done()
		ec.Subscribe("GAMES.0.ACTION", func(a action) {
			t.Logf("%+v", a)
			wait <- struct{}{}
		})
		Run(ctx)
		newGameData := struct {
			Name string
			Pid  uint64
		}{"New Game", 0}
		var res interface{}
		a := action{"drawstorm"}
		ec.RequestWithContext(ctx, "GAMES.NEW", newGameData, res)
		ec.Publish("GAMES.0.ACTION", a)
		select {
		case <-wait:
			t.Log("the function was called")
		case <-ctx.Done():
			t.Fatal("function was not called")
		}
	})
}
