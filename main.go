package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame()
	sd := &g.StormDeck
	sdis := &g.StormDiscard
	gd := &g.GearDeck

	g.DrawStormCard()
	g.DrawStormCard()
	g.DrawStormCard()
	g.DrawStormCard()
	g.DrawStormCard()

	fmt.Println("STORM DECK\n==========")
	for i := range sd.Cards {
		fmt.Println(sd.Cards[i])
	}

	fmt.Println("\nSTORM DISCARDS\n==========")
	for i := range sdis.Cards {
		fmt.Println(sdis.Cards[i])
	}

	fmt.Println("\nGEAR DECK\n==========")
	for i := range gd.Cards {
		fmt.Println(gd.Cards[i])
	}
}
