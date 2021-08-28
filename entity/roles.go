package entity

import "fmt"

type IRole interface {	
	Accuse(playerId string, g *Game)
}

type SRole struct {
	Name string	
}

func (r SRole) Accuse(playerId string, g *Game) {
	g.AccusedPlayersIds = append(g.AccusedPlayersIds, playerId)
	fmt.Printf("player with id %v has been accused", playerId)
}

type Villager struct {
	SRole
}

func CreateVillager() Villager {
	return Villager {
		SRole{Name: "villager",},
	}
}

type Werewolf struct {
	SRole	
}

func CreateWerewolf() Werewolf {
	return Werewolf {
		SRole{Name: "werewolf",},
	}
}

type Seer struct {
	SRole
}


func CreateSeer() Seer {
	return Seer {
		SRole{Name: "seer",},
	}
}

type Healer struct {
	SRole
	
}


func CreateHealer() Healer {
	return Healer {
		SRole{Name: "healer",},
	}
}




// will need to adjust this to handle games of 8+
// should it be moved to util?
func CreateRolesSlice(n int) []IRole {
	roles := []IRole{ CreateWerewolf(), CreateWerewolf(), CreateHealer(), CreateSeer() }
	n = n-4
	i := 0
	for (i < n) {
		roles = append(roles, CreateVillager())
		i++
	}
	return roles
}
