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
	d.Cards = make([]card, 0, 10)
	d.Cards = append(d.Cards,
		card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: "North",
				Distance:  1,
			},
		},
		card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: "East",
				Distance:  1,
			},
		},
		card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: "West",
				Distance:  1,
			},
		},
		card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: "South",
				Distance:  1,
			},
		},
		card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: "North",
				Distance:  2,
			},
		},
		card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: "East",
				Distance:  2,
			},
		},
		card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: "West",
				Distance:  2,
			},
		},
		card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: "South",
				Distance:  2,
			},
		},
	)
	for i := 0; i < 4; i++ {
		d.Cards = append(d.Cards, card{CardType: "STORM_PICKS_UP"})
	}
	for i := 0; i < 4; i++ {
		d.Cards = append(d.Cards, card{CardType: "SUN_BEATS_DOWN"})
	}
	return d
}
