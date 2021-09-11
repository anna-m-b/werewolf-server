package entity

var LiveGames = []*Game{}

func AddGameToLiveGames(g *Game) {
	LiveGames = append(LiveGames, g)
}

