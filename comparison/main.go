package main

import (
	"encoding/json"
	"fmt"

	"github.com/educlos/blackjack/cards"
	"github.com/educlos/blackjack/players"
)

var numberOfRounds = 10000
var numberOfDeck = 4

func main() {

	type counts struct {
		Victories int `json:"victories"`
		Loss      int `json:"loss"`
		Ties      int `json:"ties"`
	}
	var stats struct {
		BankAlike   counts `json:"bank-alike"`
		Random      counts `json:"random"`
		SmartRandom counts `json:"smart-random"`
		Basic       counts `json:"basic"`
	}

	for i := 0; i < numberOfRounds; i++ {
		var playersArray []players.Playable

		bankAlikePlayer := players.NewBankAlike("bankAlike")
		playersArray = append(playersArray, &bankAlikePlayer)

		randomPlayer := players.NewRandomPlayer("random")
		playersArray = append(playersArray, &randomPlayer)

		smartRandomPlayer := players.NewSmartRandomPlayer("smartRandom")
		playersArray = append(playersArray, &smartRandomPlayer)

		basicPlayer := players.NewBasic("basicPlayer")
		playersArray = append(playersArray, &basicPlayer)

		bank := players.NewBank()

		d := cards.GetNewDeck(numberOfDeck)
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
		if bankVal > 21 {
			fmt.Println("Bank lost")
		}
		for _, p := range playersArray {
			pVal := p.GetHandValue()
			if pVal > 21 || pVal < bankVal && bankVal <= 21 {
				switch p.GetName() {
				case "bankAlike":
					stats.BankAlike.Loss = stats.BankAlike.Loss + 1
				case "random":
					stats.Random.Loss = stats.Random.Loss + 1
				case "smartRandom":
					stats.SmartRandom.Loss = stats.SmartRandom.Loss + 1
				case "basicPlayer":
					stats.Basic.Loss = stats.Basic.Loss + 1
				}
			} else if pVal == bankVal {
				switch p.GetName() {
				case "bankAlike":
					stats.BankAlike.Ties = stats.BankAlike.Ties + 1
				case "random":
					stats.Random.Ties = stats.Random.Ties + 1
				case "smartRandom":
					stats.SmartRandom.Ties = stats.SmartRandom.Ties + 1
				case "basicPlayer":
					stats.Basic.Ties = stats.Basic.Ties + 1
				}
			} else {
				switch p.GetName() {
				case "bankAlike":
					stats.BankAlike.Victories = stats.BankAlike.Victories + 1
				case "random":
					stats.Random.Victories = stats.Random.Victories + 1
				case "smartRandom":
					stats.SmartRandom.Victories = stats.SmartRandom.Victories + 1
				case "basicPlayer":
					stats.Basic.Victories = stats.Basic.Victories + 1
				}
			}
		}
	}

	fmt.Println()
	fmt.Println()
	ctn, _ := json.Marshal(stats)
	fmt.Println(string(ctn))
}
