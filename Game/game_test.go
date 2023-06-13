package game

import "testing"

func TestGetNewGame(t *testing.T) {

	g := GetNewGame(2, 1, 1)
	if g == nil {
		t.Fatal("Game should not be nil")
	}
}
