package baas

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	cards           [40]Card
	briscola        Card // this is for semplicity the last Card
	lastCardCounter int
	playedCard      Card
	players         [2]*Player
}

func NewDeck(player1 *Player, player2 *Player) *Deck {
	seeds := []rune{'A', 'B', 'C', 'D'}
	symbols := []rune{'2', '3', '4', '5', '6', '7', 'J', 'Q', 'K', 'A'}

	var deck Deck
	deck.players[0] = player1
	deck.players[1] = player2

	deck.lastCardCounter = 0
	deck.playedCard = Card{}

	var cards [40]Card
	var counter int
	for _, seed := range seeds {
		for _, symbol := range symbols {
			cards[counter] = Card{
				symbol: symbol,
				seed:   seed,
			}
			counter++
		}
	}

	// Shuffle
	for r := 1; r < 1000; r++ {
		i := rand.Intn(39)
		j := rand.Intn(39)
		tmp := cards[i]
		cards[i] = cards[j]
		cards[j] = tmp
	}

	deck.cards = cards
	deck.briscola = cards[39]

	return &deck
}

func (d *Deck) GetBriscola() Card {
	return d.briscola
}

func (d *Deck) GetWinner(card1 Card, card2 Card) (*Player, int, error) {

	if card1.symbol == card2.symbol && card1.seed == card2.seed {
		return nil, 0, NewError(1)
	}

	winnedId, totScore := getWinnerIndex(card1, card2, d.briscola)
	return d.players[winnedId], totScore, nil

}

func (d *Deck) GetCard(playerId string) error {
	var card Card

	playerIndex, err := d.getPlayerIndex(playerId)
	if err != nil {
		return err
	}

	player := d.players[playerIndex]

	if d.lastCardCounter == 39 {
		return NewError(2)
	}
	card = d.cards[d.lastCardCounter]
	d.lastCardCounter++

	player.cards = append(player.cards, card)

	return nil
}

func (d *Deck) PlayCard(playerId string, card Card) error {

	playerIndex, err := d.getPlayerIndex(playerId)
	if err != nil {
		return err
	}

	player := d.players[playerIndex]
	err1 := player.removeCard(card)

	if err1 != nil {
		return err1
	}

	return nil
}

func (d *Deck) String() string {
	var s string

	for i, player := range d.players {
		s += fmt.Sprintf("Player %d: %s(score: %d)\n", i+1, player.id, player.score)
	}

	for _, card := range d.cards {
		s += fmt.Sprintf("%s %s %d\n", string(card.seed), string(card.symbol), symbolToScore[card.symbol])
	}

	s += fmt.Sprintf("Briscola: %s %s %d\n", string(d.briscola.seed), string(d.briscola.symbol), symbolToScore[d.briscola.symbol])

	return s
}

func getWinnerIndex(card1 Card, card2 Card, briscola Card) (index int, totScore int) {
	var score1 int
	var score2 int

	score1 = symbolToScore[card1.symbol]
	score2 = symbolToScore[card2.symbol]

	totScore = score1 + score2

	// When both cards has 0 score, the higest symbol wins
	if score1 == 0 && score2 == 0 {
		score1 = symbolToRelativeScore[card1.symbol]
		score2 = symbolToRelativeScore[card2.symbol]
	}

	// No card is briscola
	if card1.seed != briscola.seed && card2.seed != briscola.seed {
		// First card decides seed
		if card1.seed != card2.seed {
			return 0, totScore
		}

		// If seed is the same, highest score card wins
		if score1 > score2 {
			return 0, totScore
		}
		return 1, totScore
	}
	// Both cards are briscola,
	if card1.seed == briscola.seed && card2.seed == briscola.seed {
		// Highest score card wins
		if score1 > score2 {
			return 0, totScore
		}
		return 1, totScore
	}

	// One card is briscola
	if card1.seed == briscola.seed {
		return 0, totScore
	}
	return 1, totScore
}

func (d *Deck) getPlayerIndex(playerId string) (int, error) {
	if d.players[0].id == playerId {
		return 0, nil
	}
	if d.players[1].id == playerId {
		return 1, nil
	}
	return 0, NewError(3)
}
