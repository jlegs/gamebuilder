package main

import (
	"fmt"
	"gamebuilder/internal/blackjack"
)

func main() {
	fmt.Println("Hello World")
	// card := cardgame.Card{Value: "K", Suit: "H"}
	// fmt.Println(card)
	game := blackjack.BlackJackGame{}
	// game.GetPlayerInput(os.Stdin)
	// game.Play()
	game.Initialize(&blackjack.BJPlayer{Name: "Josh"})
	// fmt.Println("Game is: ", game)
	// game.Deck.Shuffle()
	// game.Deal()
	game.Play()
	fmt.Println(game.EvaluateWinRules())
	// fmt.Println(game.Deck)
	// fmt.Println(game.Deck.Cards)
	fmt.Println("PLAYERS CARDS ARE: ")
	game.Players[0].Hand.ShowHand()
	fmt.Println(game.Players[0].Hand.CalculateBJ())
}
