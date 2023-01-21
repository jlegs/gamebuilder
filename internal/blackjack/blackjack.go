package blackjack

import (
	"bufio"
	"gamebuilder/internal/cardgame"
	"io"
)

type HasBJ cardgame.Rule

func (bj HasBJ) Condition(game cardgame.Game) bool {

	// game.Players[0].Hand.ShowHand()
	return false
}

type BJRulebook cardgame.Rulebook

type BlackJackGame struct {
	cardgame.Game
	// Rulebook cardgame.Rulebook
	// Players  []cardgame.Player
	// Dealer   cardgame.Player
}

func (bj *BlackJackGame) Initialize(players []*cardgame.Player) {
	// var Rules = BJRules
	rb := cardgame.Rulebook{
		Rules: []cardgame.RuleCondition{HasBJ{}},
	}
	bj.Rulebook = &rb

	d := cardgame.NewDeck(1)
	bj.Deck = d

	bj.Players = players

	bj.Dealer = &cardgame.Player{Name: "Dealer"}

}

func (bj *BlackJackGame) Deal() {
	// c := Card{Value: "K", Suit: "h"}
	for _, player := range bj.Players {
		c := bj.Deck.DealCard()
		player.Hand.AddCard(c)
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

func (bj *BlackJackGame) EvaluateConditions() {

}

func (bj *BlackJackGame) ApplyRules() {

}
