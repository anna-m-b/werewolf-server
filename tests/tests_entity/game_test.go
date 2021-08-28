package entity_test

import (
	//"reflect"
	"testing"

	"github.com/anna-m-b/werewolf-server/entity"
)

var mPlayer1 = entity.CreatePlayer("jonny")

var mPlayer2 = entity.CreatePlayer("tally")

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
	//set up a game with 7 players
	g := entity.CreateGame(mPlayer1)
	for i := 0; i < 6; i++ {
		g.Players = append(g.Players, entity.CreatePlayer("player"))
	}
	t.Logf("players length is: %v", len(g.Players))
	
	//call assign roles
	g.AssignRoles()
	//check all players have a role
	for _, p := range g.Players {
		if p.Role == nil {
			t.Errorf("AssignRoles incorrect: got %T, wanted a role type", p.Role)
		}
	}
	t.Log(g.Players)
	//check the number of werewolves is 2, seer 1, healer 1, villagers 3
	// test that it assigns roles in a different order every time
}

// if reflect.TypeOf(p.Role) != reflect.TypeOf(entity.CreateWerewolf()) || reflect.TypeOf(p.Role) != reflect.TypeOf(entity.CreateVillager()) || reflect.TypeOf(p.Role) != reflect.TypeOf(entity.CreateSeer()) || reflect.TypeOf(p.Role) != reflect.TypeOf(entity.CreateHealer()) {
// 	t.Errorf("Player %v has no role: want a role, got: %v", p, reflect.TypeOf(p.Role))
// }