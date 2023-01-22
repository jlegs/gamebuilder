package blackjack

import (
	"fmt"
	"gamebuilder/internal/cardgame"
	"math/rand"
)

type Card struct {
	// cardgame.Card
	Value string
	Suit  string
	Order int
}

func (c *Card) Val() string {
	return c.Value
}

func (c *Card) PlayCard() cardgame.Card {
	fmt.Println("Played a card!")
	// card := cardgame.Card(c)
	return c
}

func (c *Card) ShowCard() cardgame.Card {
	fmt.Println(c.Value, c.Suit)
	return c
}

// func (c *Card) SetValues(value string, suit string) {
// 	c.Value = value
// 	c.Suit = suit
// }

type BJHand struct {
	cardgame.Hand
	// Hand  cardgame.Hand
	Cards []cardgame.Card
}

func (h *BJHand) Busted() bool {
	return h.CalculateBJ() > 21
}

func (h *BJHand) ShowHand() {
	for _, card := range h.Cards {
		card.ShowCard()
	}
	// fmt.Println(h.Cards)
}

func (h *BJHand) CalculateBJ() int {
	val := 0
	numAces := 0
	for _, card := range h.Cards {
		value := CardValues[card.Val()]
		if card.Val() == "A" {
			numAces++
		}
		val += value
	}

	for i := 0; i < numAces; i++ {
		if val > 21 {
			val -= 10
		} else {
			break
		}
	}
	return val
}

func (h *BJHand) AddCard(c *cardgame.Card) {
	h.Cards = append(h.Cards, *c)
}

func (h *BJHand) GiveCard(c *cardgame.Card) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	return c
}

func (h *BJHand) RemoveCard(c *cardgame.Card) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	return c
}

// func (h *hand)
type deck struct {
	Cards []cardgame.Card
	// Deck
}

func (d *deck) AddCard(c cardgame.Card) {
	d.Cards = append(d.Cards, c)
}

func (d *deck) DealCard() cardgame.Card {
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return c
}

func (d *deck) RemoveCard(c cardgame.Card) {
	for i, card := range d.Cards {
		if card == c {
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
		}
	}
}

func (d *deck) Shuffle() {
	// fmt.Println("Cards BEFORE are: ")
	// d.ShowCards()
	for i := 0; i < 10; i++ {
		for _, _ = range d.Cards {
			index := rand.Intn(len(d.Cards))
			card := d.Cards[index]
			d.RemoveCard(card)
			d.Cards = append(d.Cards, card)
		}
	}
	// fmt.Println("Cards AFTER shuffle are: ")
	// d.ShowCards()
}

func (d *deck) ShowCards() {
	for _, card := range d.Cards {
		card.ShowCard()
	}
}

var CardValues = map[string]int{
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

// NewDeck takes an int as an arg, and that is the number of decks to shuffle together into this new deck.
func NewDeck(num int) *deck {

	d := deck{}

	suits := []string{"h", "c", "s", "d"}

	for _, suit := range suits {
		for k, _ := range CardValues {
			c := Card{Value: k, Suit: suit}
			d.AddCard(&c)
		}

	}

	return &d
}
