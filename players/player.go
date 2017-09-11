package players

import "github.com/Etienne42/blackjack/cards"

type Player interface {
	AddNewCard(c cards.Card)
	GetValue() int
	GetHand() string
	GetName() string
	ShouldPlay() bool
}
