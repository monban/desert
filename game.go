package desert

import (
	"fmt"
)

type Game struct {
	Id           GameId
	Name         string
	StormDeck    *Deck
	StormDiscard *Deck

	GearDeck    *Deck
	GearDiscard *Deck
}

type GameAction struct {
	Action string
}

func NewGame(id GameId, name string) *Game {
	g := &Game{
		Id:           id,
		Name:         name,
		StormDeck:    NewStormDeck(),
		GearDeck:     NewGearDeck(),
		StormDiscard: &Deck{},
		GearDiscard:  &Deck{},
	}
	g.StormDeck.Shuffle()
	g.GearDeck.Shuffle()
	return g
}

func (g *Game) HandleAction(a GameAction) error {
	switch a.Action {
	case "stormdraw":
		g.DrawStormCard()
		return nil
	case "geardraw":
		g.DrawGearCard()
		return nil
	default:
		return fmt.Errorf("Unable to handle %s action", a.Action)
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
