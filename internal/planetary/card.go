package planetary

import (
	"fmt"
	"gamebuilder/internal/cardgame"
)

type PlanetaryCard struct {
	cardgame.Card
	ID             string
	Ability        Action
	PlanetLocation string
	Resources      string

	PlayRules []cardgame.RuleCondition
}

func (c *PlanetaryCard) Name() string {
	return c.ID
}

func (c PlanetaryCard) Play() cardgame.Card {
	return c
}

func (c PlanetaryCard) Val() string {
	return c.ID
}

func (c PlanetaryCard) PlayCard() cardgame.Card {
	fmt.Println("Played a card!")
	// card := cardgame.Card(c)
	return &c
}

func (c PlanetaryCard) ShowCard() cardgame.Card {
	fmt.Println(c.ID, c.Ability)
	return c
}

type Action interface {
	Act(interface{})
	Description()
	Name() string
}
type CardAbility struct {
	Action
}

type DrawFour CardAbility

func (df DrawFour) Name() string {
	return "DrawFour"
}

func (df DrawFour) Act(obj interface{}) {
	switch t := obj.(type) {
	case *PlanetaryDeck:
		fmt.Println("Called drawfour.Act()!")
		// t.Draw(4)
		c1 := t.DealCard()
		c2 := t.DealCard()
		c3 := t.DealCard()
		c4 := t.DealCard()
		fmt.Println(c1.Name())
		switch at := c1.Ability.(type) {
		case DrawFour:
			fmt.Println(at.Name())
		default:
			fmt.Println("Reached the default in the switch case")
			fmt.Println(at)
		}
		fmt.Println(c2)
		fmt.Println(c3)
		fmt.Println(c4)

	}
}

func (df DrawFour) Description() {
	fmt.Println("Playing this card will allow you to draw four cards from the deck instantly.")
}

type PlanetResources struct {
	GoldValue     int
	SpareParts    string
	MilitaryPower int
}
