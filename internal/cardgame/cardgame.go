package cardgame

type CardGame interface {
	Play()
	Initialize(Player)
	// ShowResults()
	// GetPlayerInput(*io.Reader)
	EvaluateWinRules() string
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
	WinRules   []RuleConditioner
	TurnRules  []RuleConditioner
	SetupRules []RuleConditioner
}

type Rule struct {
	RuleConditioner
}

type RuleConditioner interface {
	Condition(game interface{}) bool
}
