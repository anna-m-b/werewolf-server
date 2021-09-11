package entity

import "fmt"

type SpecialRole interface {	
	NightAction(playerId string)
}

type Werewolf struct {		
	Name string
}

func (w Werewolf) NightAction(playerId string) {
	fmt.Printf("chose victim %v", playerId)
}

func CreateWerewolf() Werewolf {
	return Werewolf {
		Name: "werewolf",
	}
}

type Seer struct {
	Name string
}
func (s Seer) NightAction(playerId string) {
	fmt.Printf("checked player for wolfyness %v", playerId)
}

func CreateSeer() Seer {
	return Seer {
		Name: "seer",
	}
}

type Healer struct {
	Name string
}

func (h Healer) NightAction(playerId string) {
	fmt.Printf("healed player %v", playerId)
}

func CreateHealer() Healer {
	return Healer {
		Name: "healer",
	}
}




// will need to adjust this to handle games of 8+
func CreateRolesSlice(n int) []SpecialRole {
	roles := []SpecialRole{ CreateWerewolf(), CreateWerewolf(), CreateHealer(), CreateSeer() }
	n = n-4
	i := 0
	for (i < n) {
		roles = append(roles, nil)
		i++
	}
	return roles
}

func RemovePlayerRoleFromSlice(s []SpecialRole, index int) []SpecialRole {
    ret := make([]SpecialRole, 0)
    ret = append(ret, s[:index]...)
    return append(ret, s[index+1:]...)
}//https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang