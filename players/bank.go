package players

import "github.com/educlos/blackjack/cards"

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

func (p *Bank) AddNewCard(c cards.Card) {
	p.hand = append(p.hand, c)
}

func (p *Bank) GetValue() (value int) {
	aceCount := 0
	for _, c := range p.hand {
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

func (p *Bank) GetHand() string {
	out := ""
	for _, c := range p.hand {
		out += c.Get() + " "
	}

	return out
}

// To improve
func (p *Bank) ShouldPlay() bool {
	if p.GetValue() < 17 {
		return true
	}
	return false
}
