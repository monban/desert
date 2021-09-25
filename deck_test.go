package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestDraw(t *testing.T) {
	is := is.New(t)
	d := Deck{}
	d.Cards = []Card{Card{"ONE", nil}, Card{"TWO", nil}}
	is.Equal(len(d.Cards), 2)
	c, ok := d.Draw()
	is.Equal(ok, true)
	is.Equal(c.CardType, "ONE")
	is.Equal(len(d.Cards), 1)
	is.Equal(d.Cards[0].CardType, "TWO")
}
