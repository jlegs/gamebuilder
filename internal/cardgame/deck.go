package cardgame

import "fmt"

type Card interface {
	ShowCard() Card
	PlayCard() Card
	Val() string
	Play() Card
	// TradeCard(*Card)
}

type Deck interface {
	Shuffle()
	DealCard() Card
	RemoveCard(Card)
	// Order()
	ShowCards()
	// ShowDisardPile()
	AddCard(Card)
}

type Hand interface {
	ShowHand()
	AddCard(Card)
	GiveCard(*Card) *Card
	RemoveCard(*Card) *Card
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

func (c *card) Val() string {
	return c.Value
}

func (c *card) Play() Card {
	fmt.Println("Play() unimplemented")
	return c
}

func (c *card) Name() string {
	return c.Value
}

type hand struct {
	Cards []Card
}

func (h *hand) ShowHand() []Card {
	fmt.Println(h.Cards)
	return h.Cards
}

func (h *hand) AddCard(c Card) {
	h.Cards = append(h.Cards, c)
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
