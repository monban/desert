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
	fmt.Println(sd)
	//fmt.Println(string(js))
}
