package players

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/educlos/blackjack/cards"
)

type RandomPlayer struct {
	Player
}

func NewRandomPlayer(name string, walletValue int) (r RandomPlayer) {
	r.Name = name
	r.Wallet = walletValue
	return
}

func (p *RandomPlayer) Play(d *cards.Deck, currentHandValue int) {
	fmt.Printf("%s\n", p.GetName())
	played := false
	if p.shouldDoubleBet(currentHandValue) {
		p.DoubleBetIfPossible()
		played = true
	}

	for p.shouldPlay() {
		played = true
		c := d.DealNextCard()
		fmt.Printf("\ttaking a card: %s\n", c.Get())
		p.addNewCard(c)
		fmt.Printf("\tNew hand value: %d\n", p.GetHandValue())
		fmt.Printf("\tNew hand: %s\n", p.GetHand())
		fmt.Println()
	}
	if !played {
		fmt.Println("\tpassed")
	}
}

func (p *RandomPlayer) shouldPlay() bool {
	if p.GetHandValue() >= 21 {
		return false
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	if num < 5 {
		return true
	}
	return false
}

func (p *RandomPlayer) Bet(ammount int) {
	if ammount > p.Wallet {
		p.CurrentBet = p.Wallet
		p.Wallet = 0
	} else {
		p.CurrentBet += ammount
		p.Wallet -= ammount
	}
}

func (p *RandomPlayer) DoubleBetIfPossible() {
	if p.CurrentBet < p.Wallet {
		p.Bet(p.CurrentBet)
		fmt.Printf("\tDoubling bet\n")
	}
}

func (p *RandomPlayer) shouldDoubleBet(bVal int) bool {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	if num < 5 {
		return true
	}
	return false
}
