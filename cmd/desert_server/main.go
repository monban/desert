package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	server "github.com/monban/desert/server"
)

func main() {
	log.Println("Starting server")
	rand.Seed(time.Now().UnixNano())
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	if err := server.Run(ctx); err != nil {
		log.Printf("%s\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
