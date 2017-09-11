package players

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
)

type BankAlike struct {
	Name string
	Hand []cards.Card
}

func (p *BankAlike) GetName() string {
	return p.Name
}

func (p *BankAlike) GetHandValue() (value int) {
	aceCount := 0
	for _, c := range p.Hand {
		if c.Value() == "A" {
			aceCount += 1
		} else {
			value += c.GetPoints()
		}
	}

	for i := 0; i < aceCount; i++ {
		tmp := value + 11
		if tmp > 17 && tmp <= 21 {
			value += 11
		} else {
			value += 1
		}
	}

	return
}

func (p *BankAlike) GetHand() string {
	out := ""
	for _, c := range p.Hand {
		out += c.Get() + " "
	}

	return out
}

func (p *BankAlike) Init(d *cards.Deck) {
	for i := 0; i < 2; i++ {
		c := d.DealNextCard()
		p.addNewCard(c)
	}
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

func (p *BankAlike) addNewCard(c cards.Card) {
	p.Hand = append(p.Hand, c)
}

func (p *BankAlike) shouldPlay() bool {
	if p.GetHandValue() < 17 {
		return true
	}
	return false
}
