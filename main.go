package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	game "github.com/jatin-malik/battleship-project/Game"
)

func main() {
	// fmt.Println("Start it.")

	// Read the input file
	file, err := os.Open("input.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	M, _ := strconv.Atoi(scanner.Text())

	// fmt.Println("Got M", M)

	scanner.Scan()
	S, _ := strconv.Atoi(scanner.Text())

	// fmt.Println("got S", S)

	scanner.Scan()
	ships1 := scanner.Text()

	scanner.Scan()
	ships2 := scanner.Text()

	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	// fmt.Println("got T", T)

	scanner.Scan()
	missiles1 := scanner.Text()

	scanner.Scan()
	missiles2 := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Make a new game
	game := game.GetNewGame(M, S, T)

	game.Player1.PlaceShips(ships1)
	game.Player2.PlaceShips(ships2)

	h1, h2 := game.FireMissiles(missiles1, missiles2)

	// fmt.Printf("Player 1 hit %d times and player2 hit %d times\n", h1, h2)

	// Check the state of the game now
	fmt.Println("Player1")
	for _, r := range game.Player1.Battleground {
		fmt.Println(r)
	}
	fmt.Println()
	fmt.Println("Player2")
	for _, r := range game.Player2.Battleground {
		fmt.Println(r)
	}
	fmt.Println()
	fmt.Println("P1:", h1)
	fmt.Println("P2:", h2)
	if h1 > h2 {
		fmt.Println("Player1 wins")
	} else if h2 > h1 {
		fmt.Println("Player2 wins")
	} else {
		fmt.Println("It is a draw")
	}

}
