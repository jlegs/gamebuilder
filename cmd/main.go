package main

import (
	"fmt"
	"gamebuilder/internal/blackjack"
	"gamebuilder/internal/cardgame"
)

func main() {
	fmt.Println("Hello World")
	// card := cardgame.Card{Value: "K", Suit: "H"}
	// fmt.Println(card)
	game := blackjack.BlackJackGame{}
	// game.GetPlayerInput(os.Stdin)
	// game.Play()
	game.Initialize([]*cardgame.Player{&cardgame.Player{Name: "Josh"}})
	// fmt.Println("Game is: ", game)
	game.Play()
	fmt.Println(game.Deck)
	game.Deal()
	// fmt.Println(game.Deck.Cards)
	fmt.Println("PLAYERS CARDS ARE: ", game.Players[0].Hand.ShowHand())
}
