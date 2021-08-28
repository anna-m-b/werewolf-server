package entity

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/anna-m-b/werewolf-server/util"
)

// do we need this?
var gameStages = [6]string{"lobby", "ready", "day", "voting", "night", "end"}

type Game struct {
	CreatedOn time.Time
	IsActive  bool
	Id        string
	Players   []Player
	Stage     string
	AccusedPlayersIds []string
	SecondedPlayerId string
	VotesToKill int
	VotesToSave int
} 

func CreateGame(p Player) Game {
	return Game{
		CreatedOn: time.Now(),
		IsActive:  true,
		Id: util.GenerateShortId(),
		Players: []Player{
			p,
		},
		Stage: gameStages[0],
		AccusedPlayersIds: make([]string, 0),
		SecondedPlayerId: "",
		VotesToKill: 0,
		VotesToSave: 0,
	}
}

func (g *Game) AddPlayer(p Player) {
	g.Players = append(g.Players, p)
}

// update this to not need input and use gameStages array + game state
func (g *Game) UpdateStage() {
	// shouldGameEnd? if so g.stage = 5 and exit this function
	// consider refactoring so we check the index it is at and increase by one
	switch g.Stage {
	case gameStages[0]:
		g.Stage = gameStages[1]
	case gameStages[1]:
		g.Stage = gameStages[2]
	case gameStages[2]:
		g.Stage = gameStages[3]
	case gameStages[3]:
		g.Stage = gameStages[4]
	}
}

func (g *Game) ToggleIsActive() {
	if g.IsActive {
		g.IsActive = false
	} else {
		g.IsActive = true
	}
}

func (g *Game) AssignRoles() {
	rs := CreateRolesSlice(len(g.Players))

	for _, p := range g.Players {
		index := rand.Intn(len(rs))
		p.Role = rs[index]
		fmt.Println(p)
		rs = RemovePlayerRoleFromSlice(rs, index)
	}
}

// func ShouldGameEnd 
// checks if enough players are dead for one side to win