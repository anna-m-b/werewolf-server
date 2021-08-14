package entity_test

import (
	"testing"

	"github.com/anna-m-b/werewolf-server/entity"
)

var mPlayer1 = entity.Player{
	UserName: "jonny",
}

func TestCreateGame(t *testing.T) {
	g := entity.CreateGame(mPlayer1)

    if g.IsActive != true {
        t.Errorf("Game IsActive was incorrect, got: %v, want: %v.", g.IsActive, true)
    }

	if g.Stage != "lobby" {
		t.Errorf("Game Stage was incorrect, got: %v, want: %v.", g.Stage, "lobby")
	}

	if len(g.Players) != 1 {
		t.Errorf("Players Length was incorrect, got: %v, want: %v.", len(g.Players), 1)
	}
}

func TestAddPlayer(t *testing.T) {
	mPlayer2 := entity.Player{
		UserName: "tally",
	}
	g := entity.CreateGame(mPlayer1)

	g.AddPlayer(mPlayer2)

	if len(g.Players) != 2 {
		t.Errorf("Players Length was incorrect, got: %v, want: %v.", len(g.Players), 2)
	}
}

func TestUpdateStage(t *testing.T) {
	g := entity.CreateGame(mPlayer1)

	g.UpdateStage("new stage")

	if g.Stage != "new stage" {
		t.Errorf("Game Stage was incorrect, got %v, want %v", g.Stage, "new stage")
	}
}
