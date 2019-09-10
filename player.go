package baas

type Player struct {
	id    string
	score int
	cards []Card
}

func NewPlayer(id string) *Player {
	cards := make([]Card, 0)
	return &Player{id: id, cards: cards}
}

func (p *Player) GetPlayerCards() []Card {
	return p.cards
}

func (p *Player) Id() string {
	return p.id
}

func (p *Player) removeCard(playedCard Card) error {
	newCards := make([]Card, 0)
	for _, card := range p.GetPlayerCards() {
		if card.symbol == playedCard.symbol && card.seed == playedCard.seed {
			continue
		}
		newCards = append(newCards, card)
	}
	p.cards = newCards
	return nil
}
