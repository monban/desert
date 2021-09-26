package main

type Game struct {
	StormDeck    Deck
	StormDiscard Deck

	GearDeck    Deck
	GearDiscard Deck
}

func NewGame() Game {
	g := Game{}
	g.StormDeck = NewStormDeck()
	g.StormDeck.Shuffle()
	g.GearDeck = NewGearDeck()
	g.GearDeck.Shuffle()
	return g
}

func (g *Game) DrawStormCard() Card {
	c, _ := g.StormDeck.Draw()
	g.StormDiscard.Add(c)
	return c
}

func (g *Game) DrawGearCard() Card {
	c, _ := g.GearDeck.Draw()
	g.GearDiscard.Add(c)
	return c
}
