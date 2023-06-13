package game

import (
	"strings"
)

// Let's have a Game struct modelling the battleship

type Game struct {
	Player1 *Player
	Player2 *Player
}

func (g *Game) FireMissiles(missiles1, missiles2 string) (int, int) {
	// player 1 and player2 fire missiles at each other
	// Update corresponding battlegrounds

	// Put it into a module , this is kinda redundant!
	hits1 := 0
	hits2 := 0

	// Player 1 fired
	for _, missile := range strings.Split(missiles1, ":") {
		x := missile[0] - '0'
		y := missile[2] - '0'

		// See it it really hits
		if g.Player2.Battleground[x][y] == "B" {
			// got him
			// fmt.Println("Player 1 got a hit on ", x, y)
			g.Player2.Battleground[x][y] = "X"
			hits1++
		} else {
			// missed it
			g.Player2.Battleground[x][y] = "O"
		}
	}

	// Player 2 fired
	for _, missile := range strings.Split(missiles2, ":") {
		x := missile[0] - '0'
		y := missile[2] - '0'

		// See it it really hits
		if g.Player1.Battleground[x][y] == "B" {
			// got him
			// fmt.Println("Player 2 got a hit on ", x, y)
			g.Player1.Battleground[x][y] = "X"
			hits2++
		} else {
			// missed it
			g.Player1.Battleground[x][y] = "O"
		}
	}

	return hits1, hits2
}

func GetNewGame(M, S, T int) *Game {
	return &Game{
		Player1: GetNewPlayer(M, S, T),
		Player2: GetNewPlayer(M, S, T),
	}
}
