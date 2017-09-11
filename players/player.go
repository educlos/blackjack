package players

import "github.com/educlos/blackjack/cards"

type Player interface {
	GetHandValue() int
	GetHand() string
	GetName() string
	Init(d *cards.Deck)
	Play(d *cards.Deck)
}
