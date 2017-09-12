package players

import (
	"fmt"
	"strings"

	"github.com/educlos/blackjack/cards"
)

type Human struct {
	Player
}

func NewHumanPlayer() (h Human) {
	fmt.Println("What is your player name?")
	inputName := ""
	fmt.Scanf("%s\n", &inputName)
	h.Name = inputName
	return
}

func (p *Human) Play(d *cards.Deck) {
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

func (p *Human) shouldPlay() bool {
	if p.GetHandValue() >= 21 {
		return false
	}
	fmt.Printf("Should player %s get an other card? (current hand: %s, value: %d)\n", p.Name, p.GetHand(), p.GetHandValue())
	answer := ""
	fmt.Scanf("%s\n", &answer)
	switch strings.ToLower(answer) {
	case "y", "yes", "oui", "o":
		return true
	case "n", "no", "non":
		return false
	}

	fmt.Printf("Unknown input %s. Skipping turn\n", answer)
	return false
}
