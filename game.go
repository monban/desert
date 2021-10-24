package desert

type Game struct {
	Id           GameId
	Name         string
	StormDeck    Deck
	StormDiscard Deck

	GearDeck    Deck
	GearDiscard Deck
}

type GameAction struct {
	Action string
}

func NewGame(id GameId, name string) Game {
	g := Game{
		Id:   id,
		Name: name,
	}
	g.StormDeck = NewStormDeck()
	g.StormDeck.Shuffle()
	g.GearDeck = NewGearDeck()
	g.GearDeck.Shuffle()
	return g
}

func (g *Game) HandleAction(a GameAction) {
	switch a.Action {
	case "stormdraw":
		g.DrawStormCard()
	case "geardraw":
		g.DrawGearCard()
	}
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
