package cardgame

import "io"

type CardGame interface {
	Play()
	Initialize()
	ShowResults()
	GetPlayerInput(*io.Reader)
	EvaluateConditions()
	ApplyRules()
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
	Rules []*Rule
}

type Rule struct {
	RuleCondition
}

type RuleCondition interface {
	Condition(CardGame)
}
