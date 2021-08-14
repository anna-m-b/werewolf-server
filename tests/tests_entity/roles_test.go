package entity_test

import (
	"testing"

	"github.com/anna-m-b/werewolf-server/entity"
)

func TestCreateRolesSlice(t *testing.T) {
	rs := entity.CreateRolesSlice(7)

	if len(rs) != 7 {
		t.Errorf("Roles Slice was incorrect, got: %v, want: %v.", len(rs), 7)
	}

}