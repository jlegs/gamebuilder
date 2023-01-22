package cardgame

type CardGame interface {
	Play()
	Initialize()
	// ShowResults()
	// GetPlayerInput(*io.Reader)
	EvaluateConditions()
	ApplyRules()
	Deal()
}

type Player interface {
	GetName() string
}

type player struct {
	Name     string
	Hand     Hand
	BankRoll int
}

type Game struct {
	GameDeck Deck
	Players  []*Player
	Dealer   Player
	Rulebook *Rulebook
	// Board Board
}

type Rulebook struct {
	Rules []RuleConditioner
}

type Rule struct {
	RuleConditioner
}

type RuleConditioner interface {
	Condition(CardGame) bool
}
