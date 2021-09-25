package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestDraw(t *testing.T) {
	is := is.NewRelaxed(t)
	d := Deck{}
	d.Cards = []Card{{"ONE", nil}, {"TWO", nil}}
	is.Equal(len(d.Cards), 2)
	c, ok := d.Draw()
	is.Equal(ok, true)
	is.Equal(c.CardType, "ONE")
	is.Equal(len(d.Cards), 1)
	is.Equal(d.Cards[0].CardType, "TWO")
}

func TestAdd(t *testing.T) {
	is := is.NewRelaxed(t)
	d := Deck{}
	d.Add(Card{"ONE", nil})
	is.Equal(1, len(d.Cards))
	is.Equal("ONE", d.Cards[0].CardType)
}
