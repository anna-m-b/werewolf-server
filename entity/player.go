package entity

import "github.com/anna-m-b/werewolf-server/util"

type Player struct {
	Id string
	UserName string
	Role IRole
	IsAlive bool
}

func CreatePlayer(name string) Player {
	return Player {
		Id: util.GenerateShortId(),
		UserName: name,
		Role: new(SRole),
		IsAlive: true,
	}
}

func (p *Player) SetRole(r IRole) {
	p.Role = r
}

func ( p *Player) KillPlayer() {
	p.IsAlive = false
}

