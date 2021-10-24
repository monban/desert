package desert

import (
	"testing"

	"github.com/matryer/is"
)

func TestNewGame(t *testing.T) {
	is := is.New(t)
	gm := GameManager{}
	gm.NewGame("foo")
	is.Equal(len(gm.games), 1)
}

func TestFindGame(t *testing.T) {
	is := is.New(t)
	gm := GameManager{}
	g, _ := gm.NewGame("foo")
	is.Equal(gm.FindGame(g.Id), gm.games[0])
}
