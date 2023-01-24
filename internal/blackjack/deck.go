package blackjack

import (
	"fmt"
	"gamebuilder/internal/cardgame"
	"math/rand"
)

type Card struct {
	// cardgame.Card
	Value     string
	Suit      string
	Order     int
	PlayRules []cardgame.RuleConditioner
}

type Is10 cardgame.Rule

func (*Is10) Condition(obj interface{}) bool {
	switch t := obj.(type) {
	case *Card:
		return CardValues[t.Value] == 10
	}
	return false
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
		for i := 0; i < len(d.Cards); i++ {
			index := rand.Intn(len(d.Cards))
			card := d.Cards[index]
			d.Cards = append(d.Cards[:index], d.Cards[index+1:]...)
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
