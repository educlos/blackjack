package house

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
)

type Bank struct {
	Name          string
	Hand          []cards.Card
	shadowCard    *cards.Card
	wallet        int
	InitialWallet int
}

func NewBank(money int) *Bank {
	p := Bank{Name: "bank"}
	p.wallet = money
	p.InitialWallet = money
	return &p
}

func (p *Bank) Play(d *cards.Deck, currentHandValue int) {
	for p.shouldPlay() {
		fmt.Printf("%s\n", p.GetName())
		c := d.DealNextCard()
		fmt.Printf("\ttaking a card: %s\n", c.Get())
		p.addNewCard(c)
		fmt.Printf("\tNew hand value: %d\n", p.GetHandValue())
		fmt.Printf("\tNew hand: %s\n", p.GetHandWithShadow())
		fmt.Println()
	}
}

func (p *Bank) shouldPlay() bool {
	if p.GetHandValue() < 17 {
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
		p.addNewCard(*p.shadowCard)
		p.shadowCard.SetShadowState(false)
	}
	return
}

func (p *Bank) GetHand() string {
	out := ""
	for _, c := range p.Hand {
		out += c.Get() + " "
	}
	if p.shadowCard.IsShadow() {
		out += "🂠"
	}

	return out
}

func (p *Bank) GetHandWithShadow() string {
	out := ""
	for _, c := range p.Hand {
		out += c.Get() + " "
	}

	if p.shadowCard.IsShadow() {
		out += p.shadowCard.Get()
	}
	return out
}

func (p *Bank) Init(d *cards.Deck) {
	c := d.DealNextCard()
	if len(p.Hand) == 0 {
		p.addNewCard(c)
	} else {
		c.SetShadowState(true)
		p.shadowCard = &c
	}
}

func (p *Bank) addNewCard(c cards.Card) {
	p.Hand = append(p.Hand, c)
}

func (p *Bank) NewRound() {
	p.Hand = []cards.Card{}
	p.shadowCard = nil
}

func (b *Bank) PayPlayer(howMuch int) {
	b.wallet -= howMuch
	if b.wallet < 0 {
		b.wallet = 0
	}
}

func (p *Bank) Win(howMuch int) {
	p.wallet += howMuch
}

func (p *Bank) GetWallet() int {
	return p.wallet
}

// For player interface
func (p *Bank) Bet(ammount int) {}

func (p *Bank) Lose() {}

func (p *Bank) ShowMoney() string {
	return ""
}

func (p *Bank) GetCurrentBet() int {
	return 0
}
