package players

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
)

type Bank struct {
	name string
	hand []cards.Card
}

var b *Bank

func GetBank() *Bank {
	if b != nil {
		return b
	}
	b = &Bank{name: "bank"}
	return b
}

func (b *Bank) GetName() string {
	return b.name
}

func (b *Bank) GetHandValue() (value int) {
	aceCount := 0
	for _, c := range b.hand {
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

func (b *Bank) GetHand() string {
	out := ""
	for _, c := range b.hand {
		out += c.Get() + " "
	}

	return out
}

func (b *Bank) Init(d *cards.Deck) {
	for i := 0; i < 2; i++ {
		c := d.DealNextCard()
		b.addNewCard(c)
	}
}

func (b *Bank) Play(d *cards.Deck) {
	for b.shouldPlay() {
		fmt.Printf("%s\n", b.GetName())
		c := d.DealNextCard()
		fmt.Printf("\ttaking a card: %s\n", c.Get())
		b.addNewCard(c)
		fmt.Printf("\tNew hand value: %d\n", b.GetHandValue())
		fmt.Printf("\tNew hand: %s\n", b.GetHand())
		fmt.Println()
	}
}

func (b *Bank) shouldPlay() bool {
	if b.GetHandValue() < 17 {
		return true
	}
	return false
}

func (b *Bank) addNewCard(c cards.Card) {
	b.hand = append(b.hand, c)
}
