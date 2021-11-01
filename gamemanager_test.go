package desert

import (
	"testing"

	"github.com/matryer/is"
)

func TestNewGame(t *testing.T) {
	is := is.New(t)
	gm := GameManager{}
	ngd := NewGameData{Name: "Foo"}
	g, err := gm.NewGame(ngd)
	is.NoErr(err)
	is.Equal(len(gm.games), 1)
	is.Equal(ngd.Name, gm.games[g.Id].Name)
}

func TestFindGame(t *testing.T) {
	is := is.New(t)
	gm := GameManager{}
	g, _ := gm.NewGame(NewGameData{"foo"})
	is.Equal(gm.FindGame(g.Id), gm.games[0])
}
