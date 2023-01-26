package blackjack

import (
	"fmt"
	"gamebuilder/internal/cardgame"
)

type HasBJ cardgame.Rule

func (bj HasBJ) Condition(game interface{}) bool {
	g, ok := game.(BlackJackGame)
	if ok {
		return g.Players[0].Hand.HasBJ()
	}
	return false
}

type Busted cardgame.Rule

func (bj Busted) Condition(game interface{}) bool {
	g, ok := game.(BlackJackGame)
	if ok {
		return g.Players[0].Hand.Busted()
	}
	return false
}

type Tie cardgame.Rule

func (bj Tie) Condition(game interface{}) bool {
	g, ok := game.(BlackJackGame)
	if ok {
		if g.Players[0].Hand.HasBJ() && !g.Dealer.Hand.HasBJ() {
			return false
		} else if !g.Players[0].Hand.HasBJ() && g.Dealer.Hand.HasBJ() {
			return false
		} else if g.Players[0].Hand.CalculateBJ() == g.Dealer.Hand.CalculateBJ() {
			return true
		}
	}
	return false
}

type PlayerStrongerHand cardgame.Rule

func (PlayerStrongerHand) Condition(game interface{}) bool {
	g, ok := game.(BlackJackGame)
	if ok {
		fmt.Println("Got OK from game typecasting to BlackJackGame. returning true or false based on values: ")
		fmt.Println(g.Players[0].Hand.CalculateBJ(), g.Dealer.Hand.CalculateBJ())
		if (!g.Players[0].Hand.Busted() && !g.Dealer.Hand.Busted()) && g.Players[0].Hand.CalculateBJ() > g.Dealer.Hand.CalculateBJ() {
			return true
		} else if !g.Players[0].Hand.Busted() && g.Dealer.Hand.Busted() {
			return true
		} else {
			return false
		}
	} else {
		fmt.Println("COULD NOT CAST TO BLACKJACKGAME")
		return false
	}
	// return false
}

type BJRulebook cardgame.Rulebook
