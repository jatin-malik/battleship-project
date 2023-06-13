package game

import "testing"

func TestPlaceShips(t *testing.T) {

	// should place the ships on a battleground corectly
	ships := "1,1:2,0:2,3:3,4:4,3"

	m := 5
	s := 5
	T := 1

	player := GetNewPlayer(m, s, T)

	player.PlaceShips(ships)

	// Assert stuff
	// Incomplete stuff
	expected := [][]string{
		{},
		{},
		{},
		{},
		{},
	}
	got := player.Battleground

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if got[i][j] == expected[i][j] {
				t.Error("Not correctly placed")
			}
		}
	}

}

func TestGetNewPlayer(t *testing.T) {
	p := GetNewPlayer(2, 1, 1)
	// Can check for proper battleground initialization also
	if p == nil {
		t.Error("Should get a properly initialized player")
	}
}
