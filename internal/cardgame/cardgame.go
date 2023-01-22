package cardgame

type CardGame interface {
	Play()
	Initialize(Player)
	// ShowResults()
	// GetPlayerInput(*io.Reader)
	EvaluateConditions() string
	ApplyRules()
	Deal()
	Player
}

type Player interface {
	GetName() string
	GetBankRoll() int
	// Hand
	// GetName() string
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
