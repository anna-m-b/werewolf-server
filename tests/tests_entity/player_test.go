package entity_test

import (
	"fmt"
	"testing"

	"github.com/anna-m-b/werewolf-server/entity"
)


func TestCreatePlayer(t *testing.T) {
	p := entity.CreatePlayer("gopher")

	if p.UserName != "gopher" {
		t.Errorf("UserName was incorrect, got: %v, want: %v.", p.UserName, "gopher")
	}

	if p.IsAlive != true {
		t.Errorf("IsAlive was incorrect, got: %v, want: %v.", p.IsAlive, true)
	}
}

func TestSetRole(t *testing.T) {
	p := entity.CreatePlayer("gopher")
	expected := entity.CreateWerewolf()
	p.SetRole(expected)

	if p.SpecialRole != expected {
		t.Errorf("Player role was incorrect, got %v, want %v", p.SpecialRole, expected)
	}
}

func TestKillPlayer(t *testing.T) {
	p := entity.CreatePlayer("gopher")

	p.KillPlayer()

	if p.IsAlive != false {
		t.Errorf("Player IsAlive was incorrect, got %v, want %v", p.IsAlive, false)
	}
}

func TestAccuse(t *testing.T) {
	g := entity.CreateGame(mPlayer1)
	gp := &g
	
	mPlayer1.Accuse("3ff", gp)

	if len(g.AccusedPlayersIds) != 1 {
		t.Errorf("Accused Players has not been added to, length is: %v, want: 1", len(g.AccusedPlayersIds))
	}

	fmt.Print(g.AccusedPlayersIds)
}

func TestSecond(t *testing.T){
	g := entity.CreateGame(mPlayer1)
	gp := &g

	mPlayer1.Second("3ff", gp)

	if g.SecondedPlayerId != "3ff" {
		t.Errorf("SecondedPlayerId incorrect: want '3ff', got %v", g.SecondedPlayerId)
	}
}