package entity_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anna-m-b/werewolf-server/entity"
)


func TestCreateRolesSlice(t *testing.T) {
	rs := entity.CreateRolesSlice(7)

	if len(rs) != 7 {
		t.Errorf("Roles Slice length was incorrect, got: %v, want: %v.", len(rs), 7)
	}
}

func TestCorrectNumberOfVillagersAdded(t *testing.T) {
	rs := entity.CreateRolesSlice(7)

	numOfVillagers := 0
	
	for i := 0; i < len(rs); i++ {
		if rs[i] == nil{
			numOfVillagers++	
		}	
	}
	
	if numOfVillagers != 3 {
		t.Errorf("number of villagers in slice wrong, got: %v, want: %v.", numOfVillagers, 3)
	}

	fmt.Println("type of rs[0]:", reflect.TypeOf(rs[0]))
	fmt.Println("type of rs[1]:", reflect.TypeOf(rs[1]))
	fmt.Println("type of rs[2]:", reflect.TypeOf(rs[2]))
	fmt.Println("type of rs[3]:", reflect.TypeOf(rs[3]))
	fmt.Println("type of rs[4]:", reflect.TypeOf(rs[4]))
	fmt.Println("type of rs[5]:", reflect.TypeOf(rs[5]))
	fmt.Println("type of rs[6]:", reflect.TypeOf(rs[6]))
}