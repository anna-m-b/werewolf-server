package entity

import (
	"fmt"
	"math/rand"
)

//var gameStages = [6]string{"lobby", "deal", "day", "voting", "night", "end"}

type GamePhase interface {
	SetNext(g *Game)
	WhichPhase() string
}

// LOBBY //

type Lobby struct {
	PlayersInLobby []*Player
	ReadyToStart   bool
}

func (l Lobby) SetNext(g *Game) {
	g.Phase = &Deal{NumPlayersReady: 0, TotalPlayers: len(g.Players)}
}

func (l Lobby) WhichPhase() string {
	return "lobby"
}

func (l *Lobby) AddPlayer(p *Player) {
	l.PlayersInLobby = append(l.PlayersInLobby, p)
	l.IsReadyToStart()
}

func (l *Lobby) IsReadyToStart() {
	if len(l.PlayersInLobby) >= 7 {
		l.ReadyToStart = true
	}
}

// DEAL //

type Deal struct {
	NumPlayersReady int
	TotalPlayers    int
}

func (d Deal) SetNext(g *Game) {
	if d.TotalPlayers == d.NumPlayersReady {
		g.Phase = &Day{AccusedPlayersIds: make([]string, 0), SecondedPlayerId: ""}
	}
}

func (d Deal) WhichPhase() string {
	return "deal"
}

func (d Deal) AssignRoles(g *Game) {
	rs := CreateRolesSlice(len(g.Players))
	for _, p := range g.Players {
		index := rand.Intn(len(rs))
		p.SpecialRole = rs[index]
		rs = RemovePlayerRoleFromSlice(rs, index)
	}

}

// DAY //

type Day struct {
	AccusedPlayersIds []string
	SecondedPlayerId  string
}

func (d Day) SetNext(g *Game) {
	//if g.shouldgameend -> end
	//else -> night
}

func (d Day) WhichPhase() string {
	return "day"
}

func (d *Day) Accuse(playerId string) {
	d.AccusedPlayersIds = append(d.AccusedPlayersIds, playerId)
	fmt.Printf("player with id %v has been accused", playerId)
}

func (d *Day) Second(playerId string) {
	if d.SecondedPlayerId == "" {
		d.SecondedPlayerId = playerId
		fmt.Printf("player with id %v has been seconded", playerId)
	} else {
		fmt.Printf("a player has already been seconded, cannot second another player")
	}
}

// VOTING //

type Voting struct {
	PlayerOnTrialId string
	VotesToKill     int
	VotesToSave     int
}

func (v Voting) SetNext(g *Game) {

}

func (v Voting) WhichPhase() string {
	return "voting"
}

func (vp *Voting) Vote(v string) {
	if v == "kill" {
		vp.VotesToKill++
	}
	if v == "save" {
		vp.VotesToSave++
	} else {
		fmt.Println("error: vote not recognised")
	}
}

// NIGHT //

type Night struct {
	HealedPlayerId string
	KilledPlayerId string
	SeenPlayerId   string
}

func (n Night) SetNext(g *Game) {
	//if g.shouldgameend -> end
	//else -> day
}

func (n Night) WhichPhase() string {
	return "night"
}
