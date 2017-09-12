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

func NewRandomPlayer(name string) (r RandomPlayer) {
	r.Name = name
	return
}

func (p *RandomPlayer) Play(d *cards.Deck) {
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
