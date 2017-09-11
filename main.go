package main

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
	"github.com/educlos/blackjack/players"
)

var numberOfDeck = 3
var numberOfPlayers = 3

func main() {
	d := cards.GetNewDeck(numberOfDeck)

	var playersArray []players.Player
	for i := 0; i < numberOfPlayers; i++ {
		p := players.BankAlike{Name: fmt.Sprintf("player%d", i)}
		playersArray = append(playersArray, &p)
	}

	// human := players.NewHumanPlayer()
	// playersArray = append(playersArray, human)

	bank := players.GetBank()

	// init
	for i := 0; i < 2; i++ {
		for _, p := range playersArray {
			c := d.DealNextCard()
			p.AddNewCard(c)
		}

		if bank.GetValue() < 17 {
			c := d.DealNextCard()
			bank.AddNewCard(c)
		}
	}
	fmt.Println()

	losingPlayers := 0
	// Actual play
	for _, p := range playersArray {
		fmt.Printf("%s's hand: %s\n", p.GetName(), p.GetHand())
	}
	fmt.Printf("bank's hand: %s\n", bank.GetHand())
	fmt.Println()

	for _, p := range playersArray {
		for p.ShouldPlay() {
			fmt.Printf("%s\n", p.GetName())
			c := d.DealNextCard()
			fmt.Printf("\ttaking a card: %s\n", c.Get())
			p.AddNewCard(c)
			fmt.Printf("\tNew hand value: %d\n", p.GetValue())
			fmt.Printf("\tNew hand: %s\n", p.GetHand())
			fmt.Println()
		}
		if p.GetValue() > 21 {
			losingPlayers += 1
		}
	}

	if losingPlayers < len(playersArray) {
		for bank.ShouldPlay() {
			fmt.Printf("bank\n")
			c := d.DealNextCard()
			fmt.Printf("\ttaking a card: %s\n", c.Get())
			bank.AddNewCard(c)
			fmt.Printf("\tNew hand value: %d\n", bank.GetValue())
			fmt.Printf("\tNew hand: %s\n", bank.GetHand())
			fmt.Println()
		}
	} else {
		fmt.Printf("bank not playing, everyone lost\n")
	}

	bankVal := bank.GetValue()
	if bankVal > 21 {
		fmt.Println("Bank lost")
	}
	for _, p := range playersArray {
		pVal := p.GetValue()
		if pVal > 21 || pVal < bankVal && bankVal <= 21 {
			fmt.Printf("%s lost with %d\n", p.GetName(), p.GetValue())
		} else if pVal == bankVal {
			fmt.Printf("%s pushed with %d\n", p.GetName(), p.GetValue())
		} else {
			fmt.Printf("%s won with %d\n", p.GetName(), p.GetValue())
		}
	}
}
