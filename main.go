package main

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
	"github.com/educlos/blackjack/players"
)

var numberOfDeck = 4

func main() {
	d := cards.GetNewDeck(numberOfDeck)

	var playersArray []players.Playable

	bankAlikePlayer := players.NewBankAlike("bankAlike")
	playersArray = append(playersArray, &bankAlikePlayer)

	randomPlayer := players.NewRandomPlayer("random")
	playersArray = append(playersArray, &randomPlayer)

	smartRandomPlayer := players.NewSmartRandomPlayer("smartRandom")
	playersArray = append(playersArray, &smartRandomPlayer)

	basicPlayer := players.NewBasic("basicPlayer")
	playersArray = append(playersArray, &basicPlayer)

	// humanPlayer := players.NewHumanPlayer()
	// playersArray = append(playersArray, &humanPlayer)

	bank := players.GetBank()

	// init
	for _, p := range playersArray {
		p.Init(&d)
	}
	bank.Init(&d)

	losingPlayers := 0
	// Actual play
	for _, p := range playersArray {
		fmt.Printf("%s's hand: %s\n", p.GetName(), p.GetHand())
	}
	fmt.Printf("bank's hand: %s\n", bank.GetHand())
	fmt.Println()

	for _, p := range playersArray {
		p.Play(&d)
		if p.GetHandValue() > 21 {
			losingPlayers += 1
		}
	}

	if losingPlayers < len(playersArray) {
		bank.Play(&d)
	} else {
		fmt.Printf("bank not playing, everyone lost\n")
	}

	bankVal := bank.GetHandValue()
	fmt.Printf("\nBank's hand: %s\n", bank.GetHandWithShadow())
	if bankVal > 21 {
		fmt.Println("Bank lost")
	}
	for _, p := range playersArray {
		pVal := p.GetHandValue()
		if pVal > 21 || pVal < bankVal && bankVal <= 21 {
			fmt.Printf("%s lost with %d\n", p.GetName(), p.GetHandValue())
		} else if pVal == bankVal && pVal != 21 {
			fmt.Printf("%s pushed with %d\n", p.GetName(), p.GetHandValue())
		} else {
			fmt.Printf("%s won with %d\n", p.GetName(), p.GetHandValue())
		}
	}
}
