package entity_test

import (
	"testing"

	"github.com/anna-m-b/werewolf-server/entity"
)


func TestCreatePlayer(t *testing.T) {
	p := entity.CreatePlayer("gopher")

	if p.UserName != "gopher" {
		t.Errorf("UserName was incorrect, got: %v, want: %v.", p.UserName, "gopher")
	}
	if p.Role != "" {
		t.Errorf("Player Role was incorrect, got: %v, want: %v.", p.Role, "")
	}
	if p.IsAlive != true {
		t.Errorf("IsAlive was incorrect, got: %v, want: %v.", p.IsAlive, true)
	}
}

func TestSetRole(t *testing.T) {
	p := entity.CreatePlayer("gopher")

	p.SetRole("werewolf")

	if p.Role != "werewolf" {
		t.Errorf("Player role was incorrect, got %v, want %v", p.Role, "werewolf")
	}
}

func TestKillPlayer(t *testing.T) {
	p := entity.CreatePlayer("gopher")

	p.KillPlayer()

	if p.IsAlive != false {
		t.Errorf("Player IsAlive was incorrect, got %v, want %v", p.IsAlive, false)
	}
}