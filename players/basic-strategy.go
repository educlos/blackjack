package players

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
)

type Basic struct {
	Player
}

func NewBasic(name string) (r Basic) {
	r.Name = name
	return
}

func (p *Basic) Play(d *cards.Deck) {
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

func (p *Basic) GetFirstHandWithoutAce() (cards.Card, error) {
	for _, c := range p.Hand {
		if c.Get() != "A" {
			return c, nil
		}
	}
	return p.Hand[0], fmt.Errorf("oops, something went wrong")
}

// Based on https://www.blackjackinfo.com/blackjack-basic-strategy-engine/
func (p *Basic) shouldPlay() bool {
	pVal := p.GetHandValue()
	if pVal >= 21 {
		return false
	}
	bVal := GetBankHandValue()
	if !p.IsHandSoft {
		if pVal <= 11 {
			return true
		}
		if pVal == 12 {
			if bVal == 4 || bVal == 5 || bVal == 6 {
				return false
			}
			return true
		}
		if pVal >= 13 && pVal <= 16 {
			if bVal <= 6 {
				return false
			}
			return true
		}
		if pVal >= 17 {
			return false
		}
	} else {
		firstCard, _ := p.GetFirstHandWithoutAce()
		switch firstCard.Get() {
		case "2", "3":
			if bVal == 5 || bVal == 6 {
				return false
			}
			return true
		case "4", "5":
			if bVal == 4 || bVal == 5 || bVal == 6 {
				return false
			}
			return true
		case "6":
			if bVal == 3 || bVal == 4 || bVal == 5 || bVal == 6 {
				return false
			}
			return true
		case "7":
			if bVal <= 8 {
				return false
			}
			return true
		case "8", "9":
			return false
		}
	}
	return false
}
