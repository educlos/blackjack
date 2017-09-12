package players

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
)

type BankAlike struct {
	Player
}

func NewBankAlike(name string) (b BankAlike) {
	b.Name = name
	return
}

func (p *BankAlike) Play(d *cards.Deck) {
	for p.shouldPlay() {
		fmt.Printf("%s\n", p.GetName())
		c := d.DealNextCard()
		fmt.Printf("\ttaking a card: %s\n", c.Get())
		p.addNewCard(c)
		fmt.Printf("\tNew hand value: %d\n", p.GetHandValue())
		fmt.Printf("\tNew hand: %s\n", p.GetHand())
		fmt.Println()
	}
}

func (p *BankAlike) shouldPlay() bool {
	if p.GetHandValue() < 17 {
		return true
	}
	return false
}
