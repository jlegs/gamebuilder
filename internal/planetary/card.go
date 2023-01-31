package planetary

import (
	"fmt"
	"gamebuilder/internal/cardgame"
)

type MilitaryCard struct {
	cardgame.Card
	ID             string
	Ability        Action
	PlanetLocation string
	Resources      PlanetResources

	PlayRules []cardgame.RuleCondition
}

func (c *MilitaryCard) Name() string {
	return c.ID
}

func (c MilitaryCard) Play() cardgame.Card {
	return &c
}

func (c MilitaryCard) Val() string {
	return c.ID
}

func (c MilitaryCard) PlayCard() cardgame.Card {
	fmt.Println("Played a card!")
	// card := cardgame.Card(c)
	return c
}

func (c MilitaryCard) ShowCard() cardgame.Card {
	fmt.Println(c.ID, c.Ability)
	return &c
}

type PlanetaryCard struct {
	cardgame.Card
	ID             string
	Ability        Action
	PlanetLocation string
	Resources      PlanetResources

	PlayRules []cardgame.RuleCondition
}

func (c *PlanetaryCard) Name() string {
	return c.ID
}

func (c PlanetaryCard) Play() cardgame.Card {
	return &c
}

func (c PlanetaryCard) Val() string {
	return c.ID
}

func (c PlanetaryCard) PlayCard() cardgame.Card {
	fmt.Println("Played a card!")
	// card := cardgame.Card(c)
	return c
}

func (c PlanetaryCard) ShowCard() cardgame.Card {
	// fmt.Println(c.ID, c.Ability)
	return &c
}

type Action interface {
	Act(interface{})
	Description()
	Name() string
}

type CardAbility struct {
	Action
}

type PlanetResources struct {
	GoldValue     int
	SpareParts    string
	MilitaryPower int
}
