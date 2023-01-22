package blackjack

import (
	"bufio"
	"gamebuilder/internal/cardgame"
	"io"
)

type BlackJackGame struct {
	// Deck cardgame.Deck
	cardgame.CardGame
	Rulebook *cardgame.Rulebook
	MinBet   int
	MaxBet   int
	Players  []*BJPlayer
	Dealer   *BJPlayer
	Deck     cardgame.Deck
	// SideGame cardgame.Game
}

func (bj *BlackJackGame) Initialize(players []*BJPlayer) {
	// var Rules = BJRules
	rb := cardgame.Rulebook{
		Rules: []cardgame.RuleConditioner{HasBJ{}},
	}
	bj.Rulebook = &rb

	d := NewDeck(1)
	bj.Deck = d

	for _, player := range players {
		newHand := BJHand{}
		player.Hand = newHand
	}

	bj.Players = players

	bj.Dealer = &BJPlayer{Name: "Dealer"}

}

func (bj *BlackJackGame) Deal() {
	// c := Card{Value: "K", Suit: "h"}
	// allPlayers := []BJPlayer{}
	for i := 0; i < 2; i++ {
		for _, player := range bj.Players {
			c := bj.Deck.DealCard()
			// fmt.Println(player)
			// fmt.Println(c)

			player.Hand.AddCard(&c)
		}
	}
}

func (bj *BlackJackGame) Play() {

}

func (bj *BlackJackGame) ReceiveBets() {
	// should this be a channel?
	// would allow us to async receive bets ....
}

func (bj *BlackJackGame) GetPlayerInput(stdin io.Reader) string {

	buf := bufio.NewScanner(stdin)
	for {
		buf.Scan()
		t := buf.Text()
		if t == "q" {
			break
		} else {
			return t
		}
	}
	return ""
}

func (bj *BlackJackGame) EvaluateConditions() string {
	for _, rule := range bj.Rulebook.Rules {
		switch t := rule.(type) {
		case HasBJ:
			if t.Condition(bj.CardGame) {
				return "IDK -- HasBJ was the rule condition that was gotten."
			}
		default:
			return "IDK"
		}
	}
	return "IDK"
}

func (bj *BlackJackGame) ApplyRules() {

}

type BJPlayer struct {
	// cardgame.Player
	Hand     BJHand
	Name     string
	BankRoll int
}

func (p *BJPlayer) GetName() string {
	return p.Name
}
