package cardgame

import "fmt"

type Card interface {
	ShowCard() Card
	PlayCard() Card
	// TradeCard(*Card)
}

type Deck interface {
	Shuffle()
	DealCard() Card
	RemoveCard(*Card)
	// Order()
	// ShowCards()
	// ShowHand()
	// ShowDisardPile()
	AddCard(Card)
}

type deck struct {
	Cards []*Card
	// Deck
}

func (d *deck) AddCard(c Card) {
	d.Cards = append(d.Cards, &c)
}

func (d *deck) DealCard() Card {
	c := card{}
	return &c
}

func (d *deck) RemoveCard(c *Card) {
	fmt.Println("RemoveCard")
}

func (d *deck) Shuffle() {
	fmt.Println("Shuffle unimplemented")
}

type card struct {
	Value string
	Suit  string
}

func (c *card) PlayCard() Card {
	return c
}

func (c *card) ShowCard() Card {
	fmt.Println(c)
	return c
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
			c := card{Value: k, Suit: suit}
			d.AddCard(&c)
		}

	}

	return &d
}
