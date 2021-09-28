package desert

import (
	"log"
)

type GameId uint64

const maxGameId GameId = ^GameId(0)

type GameManager struct {
	games map[GameId]*Game
}

func (g *GameManager) NewGame(name string) (GameId, error) {
	if g.games == nil {
		g.games = make(map[GameId]*Game)
	}
	nextId := g.nextGameId()
	game := NewGame()
	g.games[nextId] = &game
	log.Printf("Created new game with id %v, named %v", nextId, name)

	return nextId, nil
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
