package entity

import (
	"fmt"
	"time"
)

type Game struct {
	CreatedOn time.Time
	IsActive  bool
	Players   []Player
	Stage     string
}

func CreateGame(p Player) Game {
	return Game{
		CreatedOn: time.Now(),
		IsActive:  true,
		Players: []Player{
			p,
		},
		Stage: "lobby",
	}
}

func (g *Game) AddPlayer(p Player) {
	g.Players = append(g.Players, p)
	fmt.Printf("players: %v", g.Players)
}

func (g *Game) UpdateStage(s string) {
	g.Stage = s
}

func (g *Game) AssignRoles() {

	


	// g.Players[0].Role = roles[0]
	// g.Players[0].Role.SayHello()
	

}