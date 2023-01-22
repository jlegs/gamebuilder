package blackjack

import "gamebuilder/internal/cardgame"

type HasBJ cardgame.Rule

func (bj HasBJ) Condition(game cardgame.CardGame) bool {

	// game.Players[0].Hand.CalculateBJ()
	return true
}

type BJRulebook cardgame.Rulebook
