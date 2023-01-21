package blackjack

import (
	"bufio"
	"fmt"
	"gamebuilder/internal/cardgame"
	"io"
)

type HasBJ cardgame.Rule

func (bj HasBJ) Condition(game cardgame.Game) bool {

	// game.Players[0].Hand.CalculateBJ()
	return true
}

type BJRulebook cardgame.Rulebook

type BlackJackGame struct {
	Deck cardgame.Deck
	cardgame.Game
	MinBet  int
	MaxBet  int
	Players []*BJPlayer
	// Deck BJDeck
	// SideGame cardgame.Game
}

func (bj *BlackJackGame) Initialize(players []*BJPlayer) {
	// var Rules = BJRules
	rb := cardgame.Rulebook{
		Rules: []cardgame.RuleConditioner{HasBJ{}},
	}
	bj.Rulebook = &rb

	d := NewDeck(1)
	bj.GameDeck = d

	// for _, player := range players {
	// h := Hand{}
	// player.Hand = &h
	// }
	bj.Players = players

	bj.Dealer = &cardgame.Player{Name: "Dealer"}

}

func (bj *BlackJackGame) Deal() {
	// c := Card{Value: "K", Suit: "h"}
	for _, player := range bj.Players {
		c := bj.Game.GameDeck.DealCard()
		player.Hand.AddCard(c)
	}
}

func (bj *BlackJackGame) Play() {
	for _, rule := range bj.Rulebook.Rules {
		if rule.Condition(bj.Game) {
			fmt.Println("IDK")
		}
	}

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

type BJPlayer struct {
	cardgame.Player
	Hand
	Name     string
	BankRoll int
}

// type BJHand struct {
// 	// cardgame.Hand
// 	// Cards []*cardgame.Card
// 	Hand
// }
