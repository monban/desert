package desert

import (
	"testing"

	"github.com/matryer/is"
)

func TestDrawStormCard(t *testing.T) {
	is := is.NewRelaxed(t)
	g := NewGame(0, "foo")
	ol := len(g.StormDeck.cards)
	c := g.DrawStormCard()
	is.Equal(c, g.StormDiscard.cards[0])
	is.Equal(ol-1, len(g.StormDeck.cards))
}

func TestDrawGearCard(t *testing.T) {
	is := is.NewRelaxed(t)
	g := NewGame(0, "foo")
	ol := len(g.GearDeck.cards)
	c := g.DrawGearCard()
	is.Equal(c, g.GearDiscard.cards[0])
	is.Equal(ol-1, len(g.GearDeck.cards))
}
