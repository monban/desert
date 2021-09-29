package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/monban/desert"
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
	gm := &desert.GameManager{}
	opts := nats.Options{
		Name: "desert",
	}
	nc, err := opts.Connect()
	if err != nil {
		return err
	}
	go handleNewGames(ctx, nc, gm)
	log.Println("Draining connection")
	nc.Drain()
	return nil
}

func handleNewGames(ctx context.Context, nc *nats.Conn, gm *desert.GameManager) {
	ch := make(chan *nats.Msg)
	nc.ChanSubscribe("GAMES.NEW", ch)
	for {
		select {
		case msg := <-ch:
			id, _ := gm.NewGame(string(msg.Data))
			msg.Respond([]byte(fmt.Sprintf("%+v", id)))
		case <-ctx.Done():
			log.Println(ctx.Err())
		}
	}
}

func handleCardDraws(ctx context.Context, nc *nats.Conn, gm *desert.GameManager) {
	ch := make(chan *nats.Msg)
	nc.ChanSubscribe("GAME.*.DRAW.>", ch)
	for {
		select {
		case msg := <-ch:
			sub := strings.Split(msg.Subject, ".")
			//msg.Respond([]byte(fmt.Sprintf("%+v", id)))
		case <-ctx.Done():
			log.Println(ctx.Err())
		}
	}
}
