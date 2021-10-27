package server

import (
	"context"
	"sync"
	"testing"
	"time"

	router "github.com/monban/nats.router"
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
		Run(ctx, nc)
		newGameData := struct {
			Name string
			Pid  uint64
		}{"New Game", 0}
		var res interface{}
		a := action{"drawstorm"}
		ec.RequestWithContext(ctx, "GAMES.NEW", newGameData, res)
		ec.Publish("GAME.0.ACTION", a)
		select {
		case <-wait:
			t.Log("the function was called")
		case <-ctx.Done():
			t.Fatal("function was not called")
		}
	})
}

func TestInfra(t *testing.T) {
	RunServer(func(nc *nats.Conn) {
		wg := sync.WaitGroup{}
		wg.Add(1)
		handler := func(ctx context.Context, msg *nats.Msg) {
			t.Logf("%+v", msg)
			wg.Done()
		}
		r := router.New(context.Background(), nc, 0)
		r.Route(">", handler)
		nc.Publish("HELLO", []byte("world!"))
		wg.Wait()
	})
}
