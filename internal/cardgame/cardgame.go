package cardgame

import (
	"fmt"
	"io"
)

type CardGame interface {
	Play()
	// Initialize()
	// ShowResults()
	GetPlayerInput(*io.Reader)
	EvaluateConditions()
	ApplyRules()
	Deal()
}

type Player struct {
	Name string
	Hand Hand
}

type Game struct {
	Deck     deck
	Players  []Player
	Dealer   Player
	Rulebook Rulebook
}

type Rulebook struct {
	Rules []RuleCondition
}

type Rule struct {
}

type RuleCondition interface {
	Condition(Game) bool
}

type Hand interface {
	ShowHand()
	AddCard(Card)
	GiveCard(*Card) *Card
	RemoveCard(*Card)
}

type hand struct {
	Cards []card
}

func (h *hand) ShowHand() {
	fmt.Println(h.Cards)
}

func (h *hand) AddCard(c card) {
	h.Cards = append(h.Cards, c)
}
