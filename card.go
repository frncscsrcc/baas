package deck

import (
	"fmt"
)

type Card struct {
	symbol rune
	seed   rune
}

func (c Card) String() string {
	return fmt.Sprintf("{%s %s %d}", string(c.seed), string(c.symbol), symbolToScore[c.symbol])
}

func NewCard(symbol rune, seed rune) Card {
	return Card{symbol, seed}
}
