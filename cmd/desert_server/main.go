package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	server "github.com/monban/desert/server"
	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultOptions.Url)
	if err != nil {
		log.Fatalf("Unable to connect to server: %s", err)
	}
	defer nc.Close()
	log.Println("Starting server")
	rand.Seed(time.Now().UnixNano())
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	if err := server.Run(ctx, nc); err != nil {
		log.Printf("%s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
