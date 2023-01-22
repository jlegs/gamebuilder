package cardgame

type CardGame interface {
	Play()
	Initialize(Player)
	// ShowResults()
	// GetPlayerInput(*io.Reader)
	EvaluateConditions() string
	ApplyRules()
	Deal()
}

type Player interface {
	GetName() string
	GetBankRoll() int
	// Hand
	// GetName() string
}

type Playerst struct {
	Name     string
	Hand     Hand
	BankRoll int
}

type Game struct {
	GameDeck Deck
	Player   Playerst
	Dealer   Playerst
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
	Condition(game interface{}) bool
}
