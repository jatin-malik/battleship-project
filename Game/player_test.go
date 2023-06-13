package game

import "testing"

func TestPlaceShips(t *testing.T) {

	// should place the ships on a battleground corectly
	// ships := "1,1:2,0:2,3:3,4:4,3"

	// GetNewPlayer()

	// ground := make([][]string, 5)

	// // Initialize battleground with 'E' byte.

	// for i := 0; i < M; i++ {
	// 	r := make([]string, M)
	// 	for j := 0; j < M; j++ {
	// 		r[j] = "_"
	// 	}
	// 	ground[i] = r
	// }

}

func TestGetNewPlayer(t *testing.T) {
	p := GetNewPlayer(2, 1, 1)
	// Can check for proper battleground initialization also
	if p == nil {
		t.Error("Should get a properly initialized player")
	}
}
