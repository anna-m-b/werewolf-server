package entity

import (
	"fmt"

	"github.com/anna-m-b/werewolf-server/util"
)

type Player struct {
	Id string
	UserName string
	SpecialRole SpecialRole
	IsAlive bool
}

func CreatePlayer(name string) *Player {
	return &Player {
		Id: util.GenerateShortId(),
		UserName: name,
		SpecialRole: nil,
		IsAlive: true,
	}
}

func (p Player) Accuse(playerId string, g *Game) {
	g.AccusedPlayersIds = append(g.AccusedPlayersIds, playerId)
	fmt.Printf("player with id %v has been accused", playerId)
}

func (p Player) Second(playerId string, g *Game) {
	//check sconded does not already hold value
	g.SecondedPlayerId = playerId
	fmt.Printf("player with id %v has been seconded", playerId)
}

// func (r SRole) Vote(voterId string, v string, g *Game) {

// 	g.AccusedPlayersIds = append(g.AccusedPlayersIds, playerId)
// 	fmt.Printf("player with id %v has been accused", playerId)
// }

//should these  be attached to the game, and not used by any players?

func (p *Player) SetRole(r SpecialRole) {
	p.SpecialRole = r
}


func ( p *Player) KillPlayer() {
	p.IsAlive = false
}

