package cardgame

type Card interface {
	ShowCard()
	PlayCard() *card
	TradeCard(*Card)
}

type Deck interface {
	Shuffle()
	DealCard() *Card
	RemoveCard(*Card)
	Order()
	ShowCards()
	ShowHand()
	ShowDisardPile()
	AddCard(*Card)
}

type Hand interface {
	ShowHand()
	AddCard()
	GiveCard(*Card) *Card
	RemoveCard(*Card)
}

type deck struct {
	Cards []Card
	Deck
}

type card struct {
	Value string
	Suit  string
}

// NewDeck takes an int as an arg, and that is the number of decks to shuffle together into this new deck.
func NewDeck(num int) *deck {
	d := deck{}
	return &d
}
