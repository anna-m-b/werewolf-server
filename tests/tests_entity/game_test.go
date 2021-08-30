package entity_test

import (
	//"reflect"
	"testing"

	"github.com/anna-m-b/werewolf-server/entity"
)

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

	if len(g.Id) == 0 {
		t.Errorf("Game id was incorrect, got: %v, want: a character string.", g.Id)
	}
}

func TestAddPlayer(t *testing.T) {
	g := entity.CreateGame(mPlayer1)

	g.AddPlayer(mPlayer2)

	if len(g.Players) != 2 {
		t.Errorf("Players Length was incorrect, got: %v, want: %v.", len(g.Players), 2)
	}
}

func TestUpdateStage(t *testing.T) {
	g := entity.CreateGame(mPlayer1)

	g.UpdateStage()

	if g.Stage != "ready" {
		t.Errorf("Game Stage was incorrect, got %v, want %v", g.Stage, "ready")
	}
}

func TestToggleIsActive(t *testing.T) {
	g := entity.CreateGame(mPlayer1)
	g.IsActive = true

	g.ToggleIsActive()

	if g.IsActive != false {
		t.Errorf("Game's IsActive was inccorrect, got %v, want %v", g.IsActive, false)
	}
}

func TestAssignRolesToSevenPlayers(t *testing.T) {
	g := entity.CreateGame(mPlayer1)
	for i := 0; i < 6; i++ {
		g.Players = append(g.Players, entity.CreatePlayer("player"))
	}
	
	g.AssignRoles()

	roleCount := struct {
		werewolves int
		seer int
		healer int
		villagers int
	}{
		0,
		0,
		0,
		0,
	}

	for _, p := range g.Players {
		switch p.SpecialRole {
		case entity.CreateWerewolf():
			roleCount.werewolves++
		case entity.CreateHealer():
			roleCount.healer++
		case entity.CreateSeer():
			roleCount.seer++
		case nil:
			roleCount.villagers++
		}
	}
	
	if roleCount.werewolves != 2 {
		t.Errorf("Wolf count wrong: got %v, wanted 2", roleCount.werewolves)
	}
	if roleCount.healer != 1 {
		t.Errorf("healer count wrong: got %v, wanted 1", roleCount.healer)
	}
	if roleCount.seer != 1 {
		t.Errorf("seer count wrong: got %v, wanted 1", roleCount.seer)
	}
	if roleCount.villagers != 3 {
		t.Errorf("villager count wrong: got %v, wanted 3", roleCount.villagers)
	}


	// test that it assigns roles in a different order every time
}
