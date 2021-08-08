package entity

type Player struct {
	UserName string
	Role string
	IsAlive bool
}

func CreatePlayer(name string) Player {
	return Player {
		UserName: name,
		Role: "",
		IsAlive: true,
	}
}
