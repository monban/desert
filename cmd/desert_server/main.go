package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	log.Println("Starting server")
	rand.Seed(time.Now().UnixNano())
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	if err := runServer(ctx); err != nil {
		log.Printf("%s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func runServer(ctx context.Context) error {
	gm := &gameManager{}
	opts := nats.Options{
		Name: "desert",
	}
	nc, err := opts.Connect()
	if err != nil {
		return err
	}
	ch := make(chan *nats.Msg)
	_, err = nc.ChanSubscribe("GAMES.NEW", ch)
	handleNewGames(ctx, ch, gm)
	log.Println("Draining connection")
	nc.Drain()
	return nil
}

func handleNewGames(ctx context.Context, c chan *nats.Msg, gm *gameManager) {
	for {
		select {
		case msg := <-c:
			id, _ := gm.newGame(string(msg.Data))
			msg.Respond([]byte(fmt.Sprintf("%+v", id)))
		case <-ctx.Done():
			log.Println(ctx.Err())
			return
		}
	}
}
