package players

import (
	"fmt"
	"strings"

	"github.com/Etienne42/blackjack/cards"
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

func (p *Human) AddNewCard(c cards.Card) {
	p.Hand = append(p.Hand, c)
}

func (p *Human) GetValue() (value int) {
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

func (p *Human) ShouldPlay() bool {
	fmt.Printf("Should player %s get an other card? (current hand: %s, value: %d)\n", p.Name, p.GetHand(), p.GetValue())
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
