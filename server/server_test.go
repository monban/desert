package server

import (
	"context"
	"testing"

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
		ctx, done := context.WithCancel(context.Background())
		defer done()
		Run(ctx)
		nc.RequestWithContext(ctx, "GAMES.NEW", []byte(""))
		msg, _ := nc.RequestWithContext(ctx, "GAMES.0.DRAW.storm", []byte(""))
		t.Log(msg)
	})
}
