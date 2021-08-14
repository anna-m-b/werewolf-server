package entity

type Player struct {
	Id int
	UserName string
	Role PlayerRole
	IsAlive bool
}

func CreatePlayer(name string) Player {
	return Player {
		UserName: name,
		IsAlive: true,
	}
}

func (p *Player) SetRole(r PlayerRole) {
	p.Role = r
}

func ( p *Player) KillPlayer() {
	p.IsAlive = false
}

