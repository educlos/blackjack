package players

import (
	"fmt"
	"strings"

	"github.com/educlos/blackjack/cards"
)

type Human struct {
	Name string
	Hand []cards.Card
}

func NewHumanPlayer() *Human {
	fmt.Println("What is your player name?")
	inputName := ""
	fmt.Scanf("%s\n", &inputName)
	h := Human{Name: inputName}
	return &h
}

func (p *Human) GetName() string {
	return p.Name
}

func (p *Human) GetHandValue() (value int) {
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

func (p *Human) GetHand() string {
	out := ""
	for _, c := range p.Hand {
		out += c.Get() + " "
	}

	return out
}

func (p *Human) Init(d *cards.Deck) {
	for i := 0; i < 2; i++ {
		c := d.DealNextCard()
		p.addNewCard(c)
	}
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

func (p *Human) addNewCard(c cards.Card) {
	p.Hand = append(p.Hand, c)
}
