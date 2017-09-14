package main

import (
	"encoding/json"
	"fmt"

	"github.com/educlos/blackjack/house"
	"github.com/educlos/blackjack/players"
)

var numberOfRounds = 10
var numberOfDeck = 6
var defaultWallet = 100
var numberOfRoundsPerDeck = 5
var numberOfRoundsMax = 1000
var bankMoney = 1000

func main() {

	type counts struct {
		Victories int `json:"victories"`
		Loss      int `json:"loss"`
	}
	var stats struct {
		BankAlike   counts `json:"bank-alike"`
		Random      counts `json:"random"`
		SmartRandom counts `json:"smart-random"`
		Basic       counts `json:"basic-strategy"`
	}

	for i := 0; i < numberOfRounds; i++ {
		table := house.NewTable(numberOfDeck, numberOfRoundsPerDeck, numberOfRoundsMax, bankMoney)

		bankAlikePlayer := players.NewBankAlike("bankAlike", defaultWallet)
		table.RegisterPlayer(&bankAlikePlayer)

		randomPlayer := players.NewRandomPlayer("random", defaultWallet)
		table.RegisterPlayer(&randomPlayer)

		smartRandomPlayer := players.NewSmartRandomPlayer("smartRandom", defaultWallet)
		table.RegisterPlayer(&smartRandomPlayer)

		basicPlayer := players.NewBasic("basicPlayer", defaultWallet)
		table.RegisterPlayer(&basicPlayer)

		// humanPlayer := players.NewHumanPlayer()
		// table.RegisterPlayer(&humanPlayer)

		table.Play()
		s := table.GetStats()
		for p, s := range s.PlayersStats {
			switch p.GetName() {
			case "bankAlike":
				stats.BankAlike.Loss += s.Loss
				stats.BankAlike.Victories += s.Victories
			case "random":
				stats.Random.Loss += s.Loss
				stats.Random.Victories += s.Victories
			case "smartRandom":
				stats.SmartRandom.Loss += s.Loss
				stats.SmartRandom.Victories += s.Victories
			case "basicPlayer":
				stats.Basic.Loss += s.Loss
				stats.Basic.Victories += s.Victories
			}
		}
	}

	fmt.Println()
	fmt.Println()
	ctn, _ := json.Marshal(stats)
	fmt.Println(string(ctn))
}
