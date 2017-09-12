package players

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
)

type Bank struct {
	Name       string
	Hand       []cards.Card
	shadowCard cards.Card
}

var b *Bank

func GetBank() *Bank {
	if b != nil {
		return b
	}
	b = &Bank{}
	b.SetName("bank")
	return b
}

func NewBank() (p Bank) {
	p.SetName("bank")
	return
}

func (b *Bank) Play(d *cards.Deck) {
	for b.shouldPlay() {
		fmt.Printf("%s\n", b.GetName())
		c := d.DealNextCard()
		fmt.Printf("\ttaking a card: %s\n", c.Get())
		b.addNewCard(c)
		fmt.Printf("\tNew hand value: %d\n", b.GetHandValue())
		fmt.Printf("\tNew hand: %s\n", b.GetHandWithShadow())
		fmt.Println()
	}
}

func (b *Bank) shouldPlay() bool {
	if b.GetHandValue() < 17 {
		return true
	}
	return false
}

func (p *Bank) GetName() string {
	return p.Name
}

func (p *Bank) SetName(name string) {
	p.Name = name
}

func (p *Bank) GetHandValue() (value int) {
	aceCount := 0
	for _, c := range p.Hand {
		if c.Value() == "A" {
			aceCount += 1
		} else {
			value += c.GetPoints()
		}
	}

	if p.shadowCard.IsShadow() {
		if p.shadowCard.Get() == "A" {
			aceCount += 1
		} else {
			value += p.shadowCard.GetPoints()
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

	if value == 21 {
		p.addNewCard(p.shadowCard)
		p.shadowCard.SetShadowState(false)
	}
	return
}

func (p *Bank) GetHand() string {
	out := ""
	for _, c := range p.Hand {
		out += c.Get() + " "
	}

	return out
}

func (b *Bank) GetHandWithShadow() string {
	out := ""
	for _, c := range b.Hand {
		out += c.Get() + " "
	}

	if b.shadowCard.IsShadow() {
		out += b.shadowCard.Get()
	}
	return out
}

func (p *Bank) Init(d *cards.Deck) {
	c := d.DealNextCard()
	p.addNewCard(c)

	c = d.DealNextCard()
	c.SetShadowState(true)
	b.shadowCard = c
}

func (p *Bank) addNewCard(c cards.Card) {
	p.Hand = append(p.Hand, c)
}
