package players

import "github.com/educlos/blackjack/cards"

type Player interface {
	AddNewCard(c cards.Card)
	GetValue() int
	GetHand() string
	GetName() string
	ShouldPlay() bool
}
