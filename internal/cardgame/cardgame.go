package cardgame

import (
	"io"
)

type CardGame interface {
	Play()
	Initialize()
	// ShowResults()
	GetPlayerInput(*io.Reader)
	EvaluateConditions()
	ApplyRules()
	Deal()
}

type Player struct {
	Name string
	Hand
	BankRoll int
}

type Game struct {
	GameDeck Deck
	Players  []*Player
	Dealer   *Player
	Rulebook *Rulebook
	// Board Board
}

type Rulebook struct {
	Rules []RuleConditioner
}

type Rule struct {
}

type RuleConditioner interface {
	Condition(Game) bool
}

type Hand interface {
	ShowHand()
	AddCard(*Card)
	GiveCard(*Card) *Card
	RemoveCard(*Card) *Card
}

// type hand struct {
// 	Cards []Card
// }

// func (h *hand) ShowHand() []Card {
// 	fmt.Println(h.Cards)
// 	return h.Cards
// }

// func (h *hand) AddCard(c Card) {
// 	h.Cards = append(h.Cards, c)
// }
