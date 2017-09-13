package players

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
)

type BankAlike struct {
	Player
}

func NewBankAlike(name string, walletValue int) (b BankAlike) {
	b.Name = name
	b.Wallet = walletValue
	return
}

func (p *BankAlike) Play(d *cards.Deck, currentHandValue int) {
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

func (p *BankAlike) Bet(ammount int) {
	if ammount > p.Wallet {
		p.CurrentBet = p.Wallet
		p.Wallet = 0
	} else {
		p.CurrentBet += ammount
		p.Wallet -= ammount
	}
}
