package entity

import "time"

type Game struct {
	Stage string
	Players []Player
	CreatedOn time.Time
}

func CreateGame(p Player) Game {
 return Game {
	 Stage: "lobby",
	 CreatedOn: time.Now(),
	 Players: []Player {
        p,
      },
 }
}