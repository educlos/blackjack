package players

import "github.com/educlos/blackjack/cards"

type Playable interface {
	GetHandValue() int
	GetHand() string
	GetName() string
	Init(d *cards.Deck)
	Play(d *cards.Deck)
}

type Player struct {
	Name string
	Hand []cards.Card
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

func (p *Player) Init(d *cards.Deck) {
	for i := 0; i < 2; i++ {
		c := d.DealNextCard()
		p.addNewCard(c)
	}
}

func (p *Player) addNewCard(c cards.Card) {
	p.Hand = append(p.Hand, c)
}
