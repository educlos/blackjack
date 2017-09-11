package players

import "github.com/educlos/blackjack/cards"

type BankAlike struct {
	Name string
	Hand []cards.Card
}

func (p *BankAlike) GetName() string {
	return p.Name
}

func (p *BankAlike) AddNewCard(c cards.Card) {
	p.Hand = append(p.Hand, c)
}

func (p *BankAlike) GetValue() (value int) {
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

func (p *BankAlike) ShouldPlay() bool {
	if p.GetValue() < 17 {
		return true
	}
	return false
}
