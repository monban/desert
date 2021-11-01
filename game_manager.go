package desert

import (
	"log"
)

type GameId uint64

const maxGameId GameId = ^GameId(0)

type GameManager struct {
	games map[GameId]*Game
}

type NewGameData struct {
	Name string `json:"name"`
}

func (g *GameManager) NewGame(ngd NewGameData) (*Game, error) {
	if g.games == nil {
		g.games = make(map[GameId]*Game)
	}
	nextId := g.nextGameId()
	game := NewGame(nextId, ngd.Name)
	g.games[nextId] = game
	log.Printf("Created new game with id %v, named %v", game.Id, game.Name)

	return game, nil
}

func (g *GameManager) FindGame(id GameId) *Game {
	if g, ok := g.games[id]; ok {
		return g
	}
	return nil
}

func (g *GameManager) nextGameId() GameId {
	var nextId GameId
	for nextId = 0; nextId < maxGameId; nextId++ {
		if _, ok := g.games[nextId]; ok == false {
			return nextId
		}
	}
	panic("Out of game ids!")
}

func (g *GameManager) ListGames() []*Game {
	a := make([]*Game, len(g.games))
	i := 0
	for _, g := range g.games {
		a[i] = g
		i++
	}
	return a
}
