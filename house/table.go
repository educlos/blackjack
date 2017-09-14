package house

import (
	"fmt"

	"github.com/educlos/blackjack/cards"
	"github.com/educlos/blackjack/players"
)

type Table struct {
	participants          []players.Playable
	bank                  *Bank
	deck                  *cards.Deck
	stats                 stats
	numberOfDeck          int
	numberOfRoundsPerDeck int
	numberOfRoundsMax     int
}

type stats struct {
	numberOfRounds     int
	whoLostWhen        map[players.Playable]int
	moneyWonByTheBank  int
	moneyLostByTheBank int
	bankBalance        int
	PlayersStats       map[players.Playable]counts
}

type counts struct {
	Victories int `json:"victories"`
	Loss      int `json:"loss"`
}

func NewTable(numberOfDeck, numberOfRoundsPerDeck, numberOfRoundsMax, bankMoney int) *Table {
	t := Table{}
	t.bank = NewBank(bankMoney)
	d := cards.GetNewDeck(numberOfDeck)
	t.numberOfDeck = numberOfDeck
	t.numberOfRoundsPerDeck = numberOfRoundsPerDeck
	t.numberOfRoundsMax = numberOfRoundsMax
	t.stats.PlayersStats = make(map[players.Playable]counts)
	t.deck = &d
	return &t
}

func (t *Table) RegisterPlayer(p players.Playable) {
	t.participants = append(t.participants, p)
}

func (t *Table) RemovePlayer(i int) {
	if i == len(t.participants)-1 {
		t.participants = t.participants[:i]
	} else {
		t.participants = append(t.participants[:i], t.participants[i+1:]...)
	}
}

func (t *Table) GetStats() stats {
	return t.stats
}

func (t *Table) Play() {
	rounds := 0
	losers := make(map[players.Playable]int)
	moneyWonByTheBank := 0
	moneyLostByTheBank := 0
	for {
		rounds++
		if len(t.participants) == 0 || rounds > t.numberOfRoundsMax || t.bank.GetWallet() == 0 {
			break
		}

		// reset part
		if rounds%t.numberOfRoundsPerDeck == 0 {
			t.ReInitializeDeck()
		}
		initArray := append(t.participants, t.bank)
		for _, p := range initArray {
			p.NewRound()
		}

		// bets
		var toRemove []int
		for i, p := range t.participants {
			if p.GetWallet() == 0 {
				toRemove = append(toRemove, i)
				losers[p] = rounds
			} else {
				p.Bet(10)
			}
		}
		for _, i := range toRemove {
			t.RemovePlayer(i)
		}
		if len(t.participants) == 0 {
			break
		}

		// init
		for i := 0; i < 2; i++ {
			for _, p := range initArray {
				p.Init(t.deck)
			}
		}

		losingPlayers := 0
		// Pay the natural blackjack, if any
		for _, p := range t.participants {
			if p.GetHandValue() == 21 {
				fmt.Printf("%s's hand: %s (natural blackjack)\n", p.GetName(), p.GetHand())
				t.bank.PayPlayer(p.GetCurrentBet() * 3)
				moneyLostByTheBank += p.GetCurrentBet() * 3
				p.Win(3)
			}
			fmt.Printf("%s's hand: %s\n", p.GetName(), p.GetHand())
		}
		fmt.Printf("bank's hand: %s\n", t.bank.GetHand())
		fmt.Println()

		// Actual play
		for _, p := range t.participants {
			p.Play(t.deck, t.bank.GetHandValue())
			if p.GetHandValue() > 21 {
				losingPlayers += 1
			}
		}

		if losingPlayers < len(t.participants) {
			t.bank.Play(t.deck, t.bank.GetHandValue())
		} else {
			fmt.Printf("bank not playing, everyone lost\n")
		}

		bankVal := t.bank.GetHandValue()
		fmt.Printf("\nBank's hand: %s\n", t.bank.GetHandWithShadow())
		if bankVal > 21 {
			fmt.Println("Bank lost")
		}
		for _, p := range t.participants {
			pVal := p.GetHandValue()
			if pVal > 21 || pVal < bankVal && bankVal <= 21 {
				fmt.Printf("%s lost with %d\n", p.GetName(), p.GetHandValue())
				t.bank.Win(p.GetCurrentBet())
				moneyWonByTheBank += p.GetCurrentBet()
				p.Lose()
			} else if pVal == bankVal && pVal != 21 {
				fmt.Printf("%s pushed with %d\n", p.GetName(), p.GetHandValue())
			} else {
				fmt.Printf("%s won with %d\n", p.GetName(), p.GetHandValue())
				t.bank.PayPlayer(p.GetCurrentBet())
				moneyLostByTheBank += p.GetCurrentBet()
				p.Win(1)
			}
			fmt.Println(p.ShowMoney())
		}
	}

	t.stats.numberOfRounds = rounds
	t.stats.whoLostWhen = losers
	t.stats.moneyLostByTheBank = moneyLostByTheBank
	t.stats.moneyWonByTheBank = moneyWonByTheBank
	t.stats.bankBalance = moneyWonByTheBank - moneyLostByTheBank
	fmt.Printf("\n\n~~~~~ Final stats ~~~~~\n\n")
	fmt.Printf("Number of rounds: %d\n", t.stats.numberOfRounds)
	fmt.Printf("Bank's has %d$ left (initial: %d$)\n", t.bank.GetWallet(), t.bank.InitialWallet)
	fmt.Printf("Bank's has won %d$, lost %d$, for a balance of %d$\n", t.stats.moneyWonByTheBank, t.stats.moneyLostByTheBank, t.stats.bankBalance)
	for p, r := range t.stats.whoLostWhen {
		fmt.Printf("%s lost at round %d\n", p.GetName(), r)
		tmp := t.stats.PlayersStats[p]
		tmp.Loss = tmp.Loss + 1
		t.stats.PlayersStats[p] = tmp
	}
	for _, p := range t.participants {
		_, ok := losers[p]
		if !ok {
			tmp := t.stats.PlayersStats[p]
			tmp.Victories = tmp.Victories + 1
			t.stats.PlayersStats[p] = tmp
		}
	}
}

func (t *Table) ReInitializeDeck() {
	d := cards.GetNewDeck(t.numberOfDeck)
	t.deck = &d
}
