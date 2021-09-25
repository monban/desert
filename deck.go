package main

import "math/rand"

type Deck struct {
	Cards []card
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

func NewStormDeck() Deck {
	d := Deck{}
	d.Cards = make([]card, 0, 31)

	directions := []string{"North", "South", "East", "West"}
	for dir := range directions {
		for i := 0; i < 3; i++ {
			d.Cards = append(d.Cards, card{
				CardType: "STORM_MOVES",
				Storm: &stormCard{
					Direction: directions[dir],
					Distance:  1,
				},
			})
		}
		for i := 0; i < 2; i++ {
			d.Cards = append(d.Cards, card{
				CardType: "STORM_MOVES",
				Storm: &stormCard{
					Direction: directions[dir],
					Distance:  2,
				},
			})
		}
		d.Cards = append(d.Cards, card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: directions[dir],
				Distance:  3,
			},
		})
	}
	for i := 0; i < 3; i++ {
		d.Cards = append(d.Cards, card{CardType: "STORM_PICKS_UP"})
	}
	for i := 0; i < 4; i++ {
		d.Cards = append(d.Cards, card{CardType: "SUN_BEATS_DOWN"})
	}
	return d
}

func NewGearDeck() Deck {
	d := Deck{}
	d.Cards = make([]card, 0, 10)
	for i := 0; i < 3; i++ {
		d.Cards = append(d.Cards, card{CardType: "DUNE_BLASTER"})
	}
	for i := 0; i < 3; i++ {
		d.Cards = append(d.Cards, card{CardType: "JET_PACK"})
	}
	for i := 0; i < 2; i++ {
		d.Cards = append(d.Cards, card{CardType: "SOLAR_SHIELD"})
	}
	for i := 0; i < 2; i++ {
		d.Cards = append(d.Cards, card{CardType: "TERRASCOPE"})
	}
	d.Cards = append(d.Cards, card{CardType: "SECRET_WATER_RESERVE"}, card{CardType: "TIME_THROTTLE"})
	return d
}
