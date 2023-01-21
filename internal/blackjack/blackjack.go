package blackjack

import (
	"bufio"
	"fmt"
	"gamebuilder/internal/cardgame"
	"io"
)

type HasBJ cardgame.Rule

func (bj HasBJ) Condition(game cardgame.Game) bool {

	game.Players[0].Hand.ShowHand()
	return false
}

type BJRulebook cardgame.Rulebook

type BlackJackGame struct {
	Game cardgame.Game
	// Rulebook cardgame.Rulebook
	// Players  []cardgame.Player
	// Dealer   cardgame.Player
}

func (bj *BlackJackGame) Initialize(players []cardgame.Player) {
	// var Rules = BJRules
	rb := cardgame.Rulebook{
		Rules: []cardgame.RuleCondition{HasBJ{}},
	}
	bj.Game.Rulebook = rb

	d := cardgame.NewDeck(2)
	bj.Game.Deck = d

	bj.Game.Players = players

	bj.Game.Dealer = cardgame.Player{Name: "Dealer"}

}

func (bj *BlackJackGame) Deal() {
	c := Card{Value: "K", Suit: "h"}
	for _, player := range bj.Game.Players {
		player.Hand.AddCard(&c)
	}
}

func (bj *BlackJackGame) Play() {

}

func (bj *BlackJackGame) GetPlayerInput(stdin io.Reader) {

	buf := bufio.NewScanner(stdin)
	for {
		buf.Scan()
		t := buf.Text()
		if t == "q" {
			break
		} else {
			fmt.Println("enter something")
		}
	}

}

func (bj *BlackJackGame) EvaluateConditions() {

}

func (bj *BlackJackGame) ApplyRules() {

}
