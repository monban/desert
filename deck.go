package main

import "math/rand"

type Deck struct {
	Cards []Card
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// Remove a card from the top of the deck, and return that card
func (d *Deck) Draw() (Card, bool) {
	if len(d.Cards) < 1 {
		return Card{}, false
	}
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return c, true
}

func NewStormDeck() Deck {
	d := Deck{}
	d.Cards = make([]Card, 0, 31)

	directions := []string{"North", "South", "East", "West"}
	for dir := range directions {
		for i := 0; i < 3; i++ {
			d.Cards = append(d.Cards, Card{
				CardType: "STORM_MOVES",
				Storm: &stormCard{
					Direction: directions[dir],
					Distance:  1,
				},
			})
		}
		for i := 0; i < 2; i++ {
			d.Cards = append(d.Cards, Card{
				CardType: "STORM_MOVES",
				Storm: &stormCard{
					Direction: directions[dir],
					Distance:  2,
				},
			})
		}
		d.Cards = append(d.Cards, Card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: directions[dir],
				Distance:  3,
			},
		})
	}
	for i := 0; i < 3; i++ {
		d.Cards = append(d.Cards, Card{CardType: "STORM_PICKS_UP"})
	}
	for i := 0; i < 4; i++ {
		d.Cards = append(d.Cards, Card{CardType: "SUN_BEATS_DOWN"})
	}
	return d
}

func NewGearDeck() Deck {
	d := Deck{}
	d.Cards = make([]Card, 0, 10)
	for i := 0; i < 3; i++ {
		d.Cards = append(d.Cards, Card{CardType: "DUNE_BLASTER"})
	}
	for i := 0; i < 3; i++ {
		d.Cards = append(d.Cards, Card{CardType: "JET_PACK"})
	}
	for i := 0; i < 2; i++ {
		d.Cards = append(d.Cards, Card{CardType: "SOLAR_SHIELD"})
	}
	for i := 0; i < 2; i++ {
		d.Cards = append(d.Cards, Card{CardType: "TERRASCOPE"})
	}
	d.Cards = append(d.Cards, Card{CardType: "SECRET_WATER_RESERVE"}, Card{CardType: "TIME_THROTTLE"})
	return d
}
