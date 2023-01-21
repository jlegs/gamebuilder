package blackjack

import (
	"bufio"
	"fmt"
	"gamebuilder/internal/cardgame"
	"io"
)

type HasBJ cardgame.Rule

func (*HasBJ) Condition(game cardgame.CardGame) {
	// game.Players[0].Hand.ShowHand()
}

type BJRulebook cardgame.Rulebook

type BlackJackGame struct {
	// cardgame.CardGame
	Rulebook cardgame.Rulebook
	Players  []cardgame.Player
	Dealer   cardgame.Player
}

type BJRules struct {
	Rules []cardgame.Rule
}

func (bj *BlackJackGame) Initialize() {
	// var Rules = BJRules
	rb := cardgame.Rulebook{
		Rules: []*cardgame.Rule{},
	}
	bj.Rulebook = rb

}

// {
// 	Rules: []Rule{HasBJ},
// }

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
