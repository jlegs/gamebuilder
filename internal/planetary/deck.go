package planetary

import (
	"fmt"
	"gamebuilder/internal/cardgame"
	"math/rand"
	"strconv"
)

type PlanetaryDeck struct {
	cardgame.Deck
	PlanetaryCards []*PlanetaryCard
	MilitaryCards  []*MilitaryCard
}

func (d *PlanetaryDeck) Draw(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("DRAWING SINGLE CARD")
		d.DealCard()
	}
}

func (d *PlanetaryDeck) AddCard(c PlanetaryCard) {
	d.PlanetaryCards = append(d.PlanetaryCards, &c)
}

func (d *PlanetaryDeck) DealCard() PlanetaryCard {

	fmt.Println("length of cards before and after")
	fmt.Println(len(d.PlanetaryCards))
	c := d.PlanetaryCards[0]
	d.PlanetaryCards = d.PlanetaryCards[1:]
	fmt.Println(len(d.PlanetaryCards))
	return *c
}

func (d *PlanetaryDeck) RemoveCard(c PlanetaryCard) {
	fmt.Println("Removing Card")
	for i, card := range d.PlanetaryCards {
		if card == &c {
			d.PlanetaryCards = append(d.PlanetaryCards[:i], d.PlanetaryCards[i+1:]...)
		}
	}
}

func (d *PlanetaryDeck) Shuffle() {
	for i := 0; i < len(d.PlanetaryCards)*3; i++ {
		randIndex := rand.Intn(len(d.PlanetaryCards) - 1)
		card := d.PlanetaryCards[randIndex]
		d.PlanetaryCards = append(d.PlanetaryCards[:randIndex], d.PlanetaryCards[randIndex+1:]...)
		d.PlanetaryCards = append(d.PlanetaryCards, card)
	}
}

type PlanetaryHand struct {
	cardgame.Hand
	Cards []cardgame.Card
}

func (h PlanetaryHand) ShowHand() {
	for _, card := range h.Cards {
		card.ShowCard()
	}
}

func (h *PlanetaryHand) AddCard(c cardgame.Card) {
	h.Cards = append(h.Cards, c)
}

func (h *PlanetaryHand) GiveCard(c *cardgame.Card) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	return c
}

func (h *PlanetaryHand) RemoveCard(i int) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	fmt.Println("INside Remove Card. Cards are: ", h.Cards)
	c := h.Cards[i]
	h.Cards = append(h.Cards[:i], h.Cards[i+1:]...)
	return &c
}

// NewDeck takes an int as an arg, and that is the number of decks to shuffle together into this new deck.
func NewDeck(num int) *PlanetaryDeck {

	d := PlanetaryDeck{}

	for i := 0; i < 10; i++ {
		c := PlanetaryCard{
			Ability:        DrawFour{},
			ID:             strconv.Itoa(i + 1),
			Resources:      PlanetResources{},
			PlanetLocation: "Location Unknown",
		}
		d.AddCard(c)

	}

	return &d
}

// NewDeck takes an int as an arg, and that is the number of decks to shuffle together into this new deck.
func NewMilitaryDeck(num int) *MilitaryDeck {

	d := MilitaryDeck{}

	for i := 0; i < 10; i++ {
		c2 := MilitaryCard{
			Ability: StealCard{},
			ID:      strconv.Itoa(i + 11),
		}
		d.AddCard(c2)
	}

	return &d
}

type MilitaryDeck struct {
	cardgame.Deck
	PlanetaryCards []*PlanetaryCard
	MilitaryCards  []*MilitaryCard
}

func (d *MilitaryDeck) Draw(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("DRAWING SINGLE CARD")
		d.DealCard()
	}
}

func (d *MilitaryDeck) AddCard(c MilitaryCard) {
	d.MilitaryCards = append(d.MilitaryCards, &c)
}

func (d *MilitaryDeck) DealCard() MilitaryCard {

	fmt.Println("length of cards before and after")
	fmt.Println(len(d.MilitaryCards))
	c := d.MilitaryCards[0]
	d.MilitaryCards = d.MilitaryCards[1:]
	fmt.Println(len(d.MilitaryCards))
	return *c
}

func (d *MilitaryDeck) RemoveCard(c MilitaryCard) {
	fmt.Println("Removing Card")
	for i, card := range d.MilitaryCards {
		if card == &c {
			d.MilitaryCards = append(d.MilitaryCards[:i], d.MilitaryCards[i+1:]...)
		}
	}
}

func (d *MilitaryDeck) Shuffle() {
	for i := 0; i < len(d.PlanetaryCards)*3; i++ {
		randIndex := rand.Intn(len(d.PlanetaryCards) - 1)
		card := d.PlanetaryCards[randIndex]
		d.PlanetaryCards = append(d.PlanetaryCards[:randIndex], d.PlanetaryCards[randIndex+1:]...)
		d.PlanetaryCards = append(d.PlanetaryCards, card)
	}
}
