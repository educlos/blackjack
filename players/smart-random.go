package players

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/educlos/blackjack/cards"
)

type SmartRandomPlayer struct {
	Player
}

func NewSmartRandomPlayer(name string, walletValue int) (r SmartRandomPlayer) {
	r.Name = name
	r.Wallet = walletValue
	return
}

func (p *SmartRandomPlayer) Play(d *cards.Deck, currentHandValue int) {
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

func (p *SmartRandomPlayer) shouldPlay() bool {
	if p.GetHandValue() >= 21 {
		return false
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	if p.GetHandValue() < 10 {
		if num < 8 {
			return true
		}
		return false
	}

	if p.GetHandValue() < 12 {
		if num < 5 {
			return true
		}
		return false
	}

	if p.GetHandValue() < 15 {
		if num < 3 {
			return true
		}
		return false
	}

	if p.GetHandValue() < 18 {
		if num < 1 {
			return true
		}
		return false
	}

	return false
}

func (p *SmartRandomPlayer) Bet(ammount int) {
	if ammount > p.Wallet {
		p.CurrentBet = p.Wallet
		p.Wallet = 0
	} else {
		p.CurrentBet += ammount
		p.Wallet -= ammount
	}
}
