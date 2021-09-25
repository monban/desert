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
	//js, _ := json.Marshal(sd)
	for i := range sd.Cards {
		fmt.Println(sd.Cards[i])
	}
	//fmt.Println(string(js))
}
