package main

import (
	"github.com/educlos/blackjack/house"
	"github.com/educlos/blackjack/players"
)

var numberOfDeck = 6
var defaultWallet = 100
var numberOfRoundsPerDeck = 5
var numberOfRoundsMax = 1000
var bankMoney = 1000

func main() {
	table := house.NewTable(numberOfDeck, numberOfRoundsPerDeck, numberOfRoundsMax, bankMoney)

	bankAlikePlayer := players.NewBankAlike("bankAlike", defaultWallet)
	table.RegisterPlayer(&bankAlikePlayer)

	randomPlayer := players.NewRandomPlayer("random", defaultWallet)
	table.RegisterPlayer(&randomPlayer)

	smartRandomPlayer := players.NewSmartRandomPlayer("smartRandom", defaultWallet)
	table.RegisterPlayer(&smartRandomPlayer)

	basicPlayer := players.NewBasic("basicPlayer", defaultWallet)
	table.RegisterPlayer(&basicPlayer)

	bp := players.NewBasic("basicPlayer2", defaultWallet)
	table.RegisterPlayer(&bp)

	// humanPlayer := players.NewHumanPlayer()
	// table.RegisterPlayer(&humanPlayer)

	table.Play()
}
