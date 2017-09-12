package cards

import (
	"math/rand"
	"time"
)

type Deck struct {
	cards []Card
}

func GetNewDeck(howMany int) Deck {
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	colors := []string{"♠", "♥", "♣", "♦"}
	var tmp []Card
	for i := 0; i < howMany; i++ {
		for _, v := range values {
			for _, c := range colors {
				tmp = append(tmp, Card{value: v, color: c})
			}
		}
	}

	rand.Seed(time.Now().UnixNano())
	for i := len(tmp) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}

	d := Deck{tmp}

	return d
}

func (d *Deck) DealNextCard() Card {
	tmp := d.cards
	c := tmp[0]
	d.cards = tmp[1 : len(tmp)-1]
	return c
}
