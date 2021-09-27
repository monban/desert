package main

import (
	"log"

	"github.com/monban/desert"
)

type GameId uint64

const maxGameId GameId = ^GameId(0)

type gameManager struct {
	games map[GameId]*desert.Game
}

func (g *gameManager) newGame(name string) (GameId, error) {
	if g.games == nil {
		g.games = make(map[GameId]*desert.Game)
	}
	nextId := g.nextGameId()
	game := desert.NewGame()
	g.games[nextId] = &game
	log.Printf("Created new game with id %v, named %v", nextId, name)

	return nextId, nil
}

func (g *gameManager) nextGameId() GameId {
	var nextId GameId
	for nextId = 0; nextId < maxGameId; nextId++ {
		if _, ok := g.games[nextId]; ok == false {
			return nextId
		}
	}
	panic("Out of game ids!")
}
