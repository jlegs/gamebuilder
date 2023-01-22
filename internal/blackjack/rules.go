package blackjack

import (
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

type BJRulebook cardgame.Rulebook
