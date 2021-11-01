package desert

import (
	"math/rand"
)

type Deck struct {
	cards    []Card
	watchers []func(DeckEvent)
}

// A DeckEvent is a notification of a change happening to a Deck
// Action will be one of "shuffle", "add", or "draw"
type DeckEvent struct {
	Action string `json:"action"`
	Card   *Card  `json:"card"`
	Length int    `json:"length"`
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	d.notifyWatchers("shuffle", nil)
}

// Remove a card from the top of the deck, and return that card
func (d *Deck) Draw() (Card, bool) {
	if len(d.cards) < 1 {
		return Card{}, false
	}
	c := d.cards[0]
	d.cards = d.cards[1:]
	d.notifyWatchers("draw", &c)
	return c, true
}

// Add a card to the deck
func (d *Deck) Add(c Card) {
	d.cards = append(d.cards, c)
	d.notifyWatchers("add", &c)
}

// Add a callback function, to be triggered when the deck changes
// The passed function should accept a DeckEvent
func (d *Deck) Watch(fn func(DeckEvent)) {
	if d.watchers == nil {
		d.watchers = make([]func(DeckEvent), 0, 1)
	}
	d.watchers = append(d.watchers, fn)
}

func (d *Deck) notifyWatchers(action string, c *Card) {
	for _, fn := range d.watchers {
		fn(DeckEvent{
			Action: action,
			Card:   c,
			Length: len(d.cards),
		})
	}
}

// Create a default Storm deck
func NewStormDeck() *Deck {
	d := &Deck{
		cards: make([]Card, 0, 31),
	}

	directions := []string{"North", "South", "East", "West"}
	for dir := range directions {
		for i := 0; i < 3; i++ {
			d.cards = append(d.cards, Card{
				CardType: "STORM_MOVES",
				Storm: &stormCard{
					Direction: directions[dir],
					Distance:  1,
				},
			})
		}
		for i := 0; i < 2; i++ {
			d.cards = append(d.cards, Card{
				CardType: "STORM_MOVES",
				Storm: &stormCard{
					Direction: directions[dir],
					Distance:  2,
				},
			})
		}
		d.cards = append(d.cards, Card{
			CardType: "STORM_MOVES",
			Storm: &stormCard{
				Direction: directions[dir],
				Distance:  3,
			},
		})
	}
	for i := 0; i < 3; i++ {
		d.cards = append(d.cards, Card{CardType: "STORM_PICKS_UP"})
	}
	for i := 0; i < 4; i++ {
		d.cards = append(d.cards, Card{CardType: "SUN_BEATS_DOWN"})
	}
	return d
}

// Create a default Gear deck
func NewGearDeck() *Deck {
	d := &Deck{
		cards: make([]Card, 0, 31),
	}
	for i := 0; i < 3; i++ {
		d.cards = append(d.cards, Card{CardType: "DUNE_BLASTER"})
	}
	for i := 0; i < 3; i++ {
		d.cards = append(d.cards, Card{CardType: "JET_PACK"})
	}
	for i := 0; i < 2; i++ {
		d.cards = append(d.cards, Card{CardType: "SOLAR_SHIELD"})
	}
	for i := 0; i < 2; i++ {
		d.cards = append(d.cards, Card{CardType: "TERRASCOPE"})
	}
	d.cards = append(d.cards, Card{CardType: "SECRET_WATER_RESERVE"}, Card{CardType: "TIME_THROTTLE"})
	return d
}
