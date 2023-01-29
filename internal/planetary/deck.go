package planetary

import (
	"fmt"
	"gamebuilder/internal/cardgame"
)

type PlanetaryDeck struct {
	cardgame.Deck
	Cards []PlanetaryCard
}

func (d *PlanetaryDeck) Draw(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("DRAWING SINGLE CARD")
		d.DealCard()
	}
}

func (d *PlanetaryDeck) AddCard(c PlanetaryCard) {
	d.Cards = append(d.Cards, c)
}

func (d *PlanetaryDeck) DealCard() *PlanetaryCard {
	c := d.Cards[0]
	d.Cards = d.Cards[1:]
	return &c
}

func (d *PlanetaryDeck) RemoveCard(c PlanetaryCard) {
	fmt.Println("RemoveCard")
}

func (d *PlanetaryDeck) Shuffle() {
	fmt.Println("Shuffle unimplemented")
}

type PlanetaryHand struct {
	cardgame.Hand
	Cards []cardgame.Card
}

func (h *PlanetaryHand) ShowHand() {
	for _, card := range h.Cards {
		card.ShowCard()
	}
}

func (h *PlanetaryHand) AddCard(c *cardgame.Card) {
	h.Cards = append(h.Cards, *c)
}

func (h *PlanetaryHand) GiveCard(c *cardgame.Card) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	return c
}

func (h *PlanetaryHand) RemoveCard(c *cardgame.Card) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	return c
}

// NewDeck takes an int as an arg, and that is the number of decks to shuffle together into this new deck.
func NewDeck(num int) *PlanetaryDeck {

	d := PlanetaryDeck{}

	// suits := []string{"h", "c", "s", "d"}

	for i := 0; i < 10; i++ {
		c := PlanetaryCard{
			Ability:        DrawFour{},
			ID:             "joweijoiej",
			Resources:      "IDK",
			PlanetLocation: "ALso idk",
		}
		d.AddCard(c)
		fmt.Println(c)

	}

	return &d
}
