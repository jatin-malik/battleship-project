package game

import (
	"strings"
)

type Player struct {
	Battleground [][]string
	Ships        int
	Missiles     int
}

func GetNewPlayer(M, S, T int) *Player {

	ground := make([][]string, M)

	// Initialize battleground with 'E' byte.

	for i := 0; i < M; i++ {
		r := make([]string, M)
		for j := 0; j < M; j++ {
			r[j] = "_"
		}
		ground[i] = r
	}

	return &Player{
		Battleground: ground,
		Ships:        S,
		Missiles:     T,
	}
}

func (p *Player) PlaceShips(ships string) {
	for _, ship := range strings.Split(ships, ":") {
		x := ship[0] - '0'
		y := ship[2] - '0'
		// fmt.Println("Put ship at coords ", x, y)
		p.Battleground[x][y] = "B"
	}
}
