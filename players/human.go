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

	fmt.Printf("What is %s's wallet value?\n", inputName)
	inputWallet := 0
	fmt.Scanf("%d\n", &inputWallet)

	h.Name = inputName
	h.Wallet = inputWallet
	return
}

func (p *Human) Play(d *cards.Deck, currentHandValue int) {
	if p.shouldDoubleBet() {
		fmt.Printf("%s\n", p.GetName())
		p.DoubleBetIfPossible()
	}
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

func (p *Human) Bet(ammount int) {
	fmt.Printf("How much does %s want to bet? (current wallet: %d$)\n", p.Name, p.Wallet)
	inputBet := 0
	fmt.Scanf("%d\n", &inputBet)
	counter := 0
	for inputBet > p.Wallet {
		if counter > 5 {
			break
		}
		fmt.Println("Bet too high. Please re-bet")
		fmt.Scanf("%d\n", &inputBet)
		counter++
	}
	if counter == 5 {
		fmt.Println("Too bad for you, you will bet everything")
		p.CurrentBet += p.Wallet
		p.Wallet = 0
	} else {
		p.CurrentBet += inputBet
		p.Wallet -= inputBet
	}
}

func (p *Human) DoubleBetIfPossible() {
	if p.CurrentBet < p.Wallet {
		p.Wallet -= p.CurrentBet
		p.CurrentBet *= 2
		fmt.Printf("\tDoubling bet\n")
	} else {
		fmt.Printf("\tDoubling bet is not allowed (not enought money left)\n")
	}
}

func (p *Human) shouldDoubleBet() bool {
	fmt.Printf("Should player %s bouble his bet? (hand: %s, current bet: %d, wallet: %d)\n", p.Name, p.GetHand(), p.CurrentBet, p.Wallet)
	answer := ""
	fmt.Scanf("%s\n", &answer)
	switch strings.ToLower(answer) {
	case "y", "yes", "oui", "o":
		return true
	case "n", "no", "non":
		return false
	}

	fmt.Printf("Unknown input %s. Not doubling\n", answer)
	return false
}
