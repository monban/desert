package desert

import (
	"testing"

	"github.com/matryer/is"
)

func TestDrawStormCard(t *testing.T) {
	is := is.NewRelaxed(t)
	g := NewGame()
	ol := len(g.StormDeck.Cards)
	c := g.DrawStormCard()
	is.Equal(c, g.StormDiscard.Cards[0])
	is.Equal(ol-1, len(g.StormDeck.Cards))
}

func TestDrawGearCard(t *testing.T) {
	is := is.NewRelaxed(t)
	g := NewGame()
	ol := len(g.GearDeck.Cards)
	c := g.DrawGearCard()
	is.Equal(c, g.GearDiscard.Cards[0])
	is.Equal(ol-1, len(g.GearDeck.Cards))
}
