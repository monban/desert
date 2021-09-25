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
	g.StormDeck = NewGearDeck()
	return g
}

func (g *Game) DrawStormCard() Card {
	c, _ := g.StormDeck.Draw()
	return c
}
