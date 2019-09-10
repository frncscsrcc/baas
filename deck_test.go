package baas

import (
	"fmt"
	"testing"
)

func TestGetWinnerIndex(t *testing.T) {
	var winnerIndex int
	var totScore int

	briscola := NewCard('2', 'A')

	type game struct {
		card1       Card
		card2       Card
		winnerIndex int
		totScore    int
	}

	games := make([]game, 0)

	// Just a briscola, briscola wins
	games = append(games, game{NewCard('2', 'A'), NewCard('2', 'B'), 0, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('3', 'B'), 0, 10})

	// Two briscolas, higest briscola wins
	games = append(games, game{NewCard('2', 'A'), NewCard('3', 'A'), 1, 10})

	// No briscolas, different seeds, first card seed wins
	games = append(games, game{NewCard('2', 'B'), NewCard('3', 'C'), 0, 10})

	// No briscolas, same seed, higest card wins
	games = append(games, game{NewCard('2', 'B'), NewCard('3', 'B'), 1, 10})
	games = append(games, game{NewCard('2', 'B'), NewCard('4', 'B'), 1, 0})

	// Check scores
	games = append(games, game{NewCard('2', 'A'), NewCard('3', 'A'), 1, 10})
	games = append(games, game{NewCard('2', 'A'), NewCard('4', 'A'), 1, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('5', 'A'), 1, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('6', 'A'), 1, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('7', 'A'), 1, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('J', 'A'), 1, 2})
	games = append(games, game{NewCard('2', 'A'), NewCard('Q', 'A'), 1, 3})
	games = append(games, game{NewCard('2', 'A'), NewCard('K', 'A'), 1, 4})
	games = append(games, game{NewCard('2', 'A'), NewCard('A', 'A'), 1, 11})

	// Check totScore
	games = append(games, game{NewCard('J', 'A'), NewCard('Q', 'A'), 1, 5})

	for _, g := range games {
		winnerIndex, totScore = getWinnerIndex(g.card1, g.card2, briscola)

		if winnerIndex != g.winnerIndex {
			t.Error(fmt.Sprintf("%v - %v : WinnerIndex is not %d", g.card1, g.card2, g.winnerIndex))
		}
		if totScore != g.totScore {
			t.Error(fmt.Sprintf("%v - %v : Score is not %d", g.card1, g.card2, g.totScore))
		}
	}
}
