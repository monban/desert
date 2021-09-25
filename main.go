package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	sd := NewStormDeck()
	sd.Shuffle()

	gd := NewGearDeck()
	gd.Shuffle()

	fmt.Println("STORM DECK\n==========")
	for i := range sd.Cards {
		fmt.Println(sd.Cards[i])
	}

	fmt.Println("\nGEAR DECK\n==========")
	for i := range gd.Cards {
		fmt.Println(gd.Cards[i])
	}
}
