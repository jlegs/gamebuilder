package blackjack

import (
	"fmt"
	"gamebuilder/internal/cardgame"
)

type Card struct {
	Value string
	Suit  string
}

func (c *Card) PlayCard() cardgame.Card {
	fmt.Println("Played a card!")
	// card := cardgame.Card(c)
	return c
}

func (c *Card) ShowCard() cardgame.Card {
	fmt.Println(c)
	return c
}
