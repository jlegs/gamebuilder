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

type Hand struct {
	// cardgame.Hand
	Cards []*cardgame.Card
}

func (h *Hand) ShowHand() {
	fmt.Println(h.Cards)
}

func (h *Hand) CalculateBJP() int {
	return 21
}
func (h *Hand) AddCard(c cardgame.Card) {
	h.Cards = append(h.Cards, &c)
}

func (h *Hand) GiveCard(c *cardgame.Card) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	return c
}

func (h *Hand) RemoveCard(c *cardgame.Card) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	return c
}

// func (h *hand)
type deck struct {
	Cards []*cardgame.Card
	// Deck
}

func (d *deck) AddCard(c cardgame.Card) {
	d.Cards = append(d.Cards, &c)
}

func (d *deck) DealCard() cardgame.Card {
	c := d.Cards[0]
	return *c
}

func (d *deck) RemoveCard(c *cardgame.Card) {
	fmt.Println("RemoveCard")
}

func (d *deck) Shuffle() {
	fmt.Println("Shuffle unimplemented")
}

// NewDeck takes an int as an arg, and that is the number of decks to shuffle together into this new deck.
func NewDeck(num int) *deck {

	d := deck{}
	cardValues := map[string]int{
		"K":  10,
		"Q":  10,
		"J":  10,
		"10": 10,
		"9":  9,
		"8":  8,
		"7":  7,
		"6":  6,
		"5":  5,
		"4":  4,
		"3":  3,
		"2":  2,
		"A":  11,
	}

	suits := []string{"h", "c", "s", "d"}

	for _, suit := range suits {
		for k, _ := range cardValues {
			c := Card{Value: k, Suit: suit}
			d.AddCard(&c)
		}

	}

	return &d
}
