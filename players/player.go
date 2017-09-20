package players

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
)

type Playable interface {
	GetHandValue() int
	GetHand() string
	GetName() string
	Init(d *cards.Deck)
	Play(*cards.Deck, int)
	Bet(int)
	DoubleBetIfPossible()
	Lose()
	Win(int)
	ShowMoney() string
	NewRound()
	GetWallet() int
	GetCurrentBet() int
}

type Player struct {
	Name       string
	Hand       []cards.Card
	IsHandSoft bool
	Wallet     int
	CurrentBet int
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) SetName(name string) {
	p.Name = name
}

func (p *Player) GetHandValue() (value int) {
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
			p.IsHandSoft = true
		} else {
			value += 1
		}
	}

	return
}

func (p *Player) GetHand() string {
	out := ""
	for _, c := range p.Hand {
		out += c.Get() + " "
	}

	return out
}

func (p *Player) GetWallet() int {
	return p.Wallet
}

func (p *Player) GetCurrentBet() int {
	return p.CurrentBet
}

func (p *Player) Init(d *cards.Deck) {
	c := d.DealNextCard()
	p.addNewCard(c)
}

func (p *Player) addNewCard(c cards.Card) {
	p.Hand = append(p.Hand, c)
}

func (p *Player) Lose() {
	p.CurrentBet = 0
}

func (p *Player) Win(ratio int) {
	p.Wallet = p.Wallet + p.CurrentBet + p.CurrentBet*ratio
	p.CurrentBet = 0
}

func (p *Player) ShowMoney() string {
	out := fmt.Sprintf("%s has %d$", p.Name, p.Wallet)
	if p.CurrentBet != 0 {
		out += fmt.Sprintf(" (and %d$ on bet)", p.CurrentBet)
	}
	return out
}

func (p *Player) NewRound() {
	p.Hand = []cards.Card{}
	p.IsHandSoft = false
}
