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
	Hand hand
}

type Game struct {
	Deck
	Players  []*Player
	Dealer   *Player
	Rulebook *Rulebook
	// Board Board
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
	Cards []Card
}

func (h *hand) ShowHand() []Card {
	fmt.Println(h.Cards)
	return h.Cards
}

func (h *hand) AddCard(c Card) {
	h.Cards = append(h.Cards, c)
}
