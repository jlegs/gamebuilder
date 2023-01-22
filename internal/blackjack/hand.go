package blackjack

import "gamebuilder/internal/cardgame"

type BJHand struct {
	cardgame.Hand
	// Hand  cardgame.Hand
	Cards []cardgame.Card
}

func (h *BJHand) Busted() bool {
	return h.CalculateBJ() > 21
}

func (h *BJHand) ShowHand() {
	for _, card := range h.Cards {
		card.ShowCard()
	}
}

func (h *BJHand) CalculateBJ() int {
	val := 0
	numAces := 0
	for _, card := range h.Cards {
		value := CardValues[card.Val()]
		if card.Val() == "A" {
			numAces++
		}
		val += value
	}

	for i := 0; i < numAces; i++ {
		if val > 21 {
			val -= 10
		} else {
			break
		}
	}
	return val
}

func (h *BJHand) AddCard(c *cardgame.Card) {
	h.Cards = append(h.Cards, *c)
}

func (h *BJHand) GiveCard(c *cardgame.Card) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	return c
}

func (h *BJHand) RemoveCard(c *cardgame.Card) *cardgame.Card {
	// h.Cards = append(h.Cards, c)
	return c
}

func (h *BJHand) HasBJ() bool {
	if h.CalculateBJ() == 21 && len(h.Cards) == 2 {
		return true
	}
	return false
}
