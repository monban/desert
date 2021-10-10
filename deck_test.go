package desert

import (
	"testing"

	"github.com/matryer/is"
)

func TestDraw(t *testing.T) {
	is := is.NewRelaxed(t)
	d := Deck{}
	d.cards = []Card{{"ONE", nil}, {"TWO", nil}}
	is.Equal(len(d.cards), 2)
	c, ok := d.Draw()
	is.Equal(ok, true)
	is.Equal(c.CardType, "ONE")
	is.Equal(len(d.cards), 1)
	is.Equal(d.cards[0].CardType, "TWO")
}

func TestAdd(t *testing.T) {
	is := is.NewRelaxed(t)
	d := Deck{}
	d.Add(Card{"ONE", nil})
	is.Equal(1, len(d.cards))
	is.Equal("ONE", d.cards[0].CardType)
}

func TestWatchDraw(t *testing.T) {
	is := is.NewRelaxed(t)
	var a DeckAction
	var sentCard Card = Card{"ONE", nil}
	fn := func(da DeckAction) {
		a = da
	}
	d := Deck{}
	d.Add(sentCard)

	d.Watch(fn)
	d.Draw()

	is.Equal(sentCard, *a.Card)
}
