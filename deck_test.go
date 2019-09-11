package baas

import (
	"fmt"
	"testing"
)

func TestFindWinnerIndex(t *testing.T) {
	briscola := NewCard('2', 'A')

	type game struct {
		card1       Card
		card2       Card
		winnerIndex int
		totScore    int
		errorCode   int
	}

	games := make([]game, 0)

	// Just a briscola, briscola wins
	games = append(games, game{NewCard('2', 'A'), NewCard('2', 'B'), 0, 0, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('3', 'B'), 0, 10, 0})

	// Two briscolas, higest briscola wins
	games = append(games, game{NewCard('2', 'A'), NewCard('3', 'A'), 1, 10, 0})

	// No briscolas, different seeds, first card seed wins
	games = append(games, game{NewCard('2', 'B'), NewCard('3', 'C'), 0, 10, 0})

	// No briscolas, same seed, higest card wins
	games = append(games, game{NewCard('2', 'B'), NewCard('3', 'B'), 1, 10, 0})
	games = append(games, game{NewCard('2', 'B'), NewCard('4', 'B'), 1, 0, 0})

	// Check scores
	games = append(games, game{NewCard('2', 'A'), NewCard('3', 'A'), 1, 10, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('4', 'A'), 1, 0, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('5', 'A'), 1, 0, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('6', 'A'), 1, 0, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('7', 'A'), 1, 0, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('J', 'A'), 1, 2, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('Q', 'A'), 1, 3, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('K', 'A'), 1, 4, 0})
	games = append(games, game{NewCard('2', 'A'), NewCard('A', 'A'), 1, 11, 0})

	// Check totScore
	games = append(games, game{NewCard('J', 'A'), NewCard('Q', 'A'), 1, 5, 0})

	// Same card error
	games = append(games, game{NewCard('2', 'A'), NewCard('2', 'A'), 0, 0, 1})

	for _, g := range games {
		winnerIndex, totScore, err := findWinner(g.card1, g.card2, briscola)
		
		var errorCode int

		if baasError, ok := err.(BaasError); ok {
		    errorCode = baasError.GetCode()
		}

		if g.errorCode > 0 {
			if g.errorCode != errorCode {
				t.Error(fmt.Sprintf("%v - %v : Expected error code %d", g.card1, g.card2, g.errorCode))
			}
			continue
		}
		if winnerIndex != g.winnerIndex {
			t.Error(fmt.Sprintf("%v - %v : WinnerIndex is not %d", g.card1, g.card2, g.winnerIndex))
		}
		if totScore != g.totScore {
			t.Error(fmt.Sprintf("%v - %v : Score is not %d", g.card1, g.card2, g.totScore))
		}

	}
}

func TestGetCard(t *testing.T) {
	d := NewDeck()
	for i := 0; i < 39; i++ {
		_, err := d.GetCard()
		if err != nil {
			t.Error(fmt.Sprintf("Not error expected: there are more card in the deck"))	
		}
	}
	_, err := d.GetCard()
	if baasError, ok := err.(BaasError); ok {
		if baasError.GetCode() != 2 {
			t.Error(fmt.Sprintf("Expected error 2 (no more card to play)"))
		}
	}else{
		t.Error(fmt.Sprintf("Expected error: no more cards to play"))
	}

}