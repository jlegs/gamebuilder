package main

import (
	"fmt"
	"gamebuilder/internal/blackjack"
	"os"
)

func main() {
	fmt.Println("Hello World")
	// card := cardgame.Card{Value: "K", Suit: "H"}
	// fmt.Println(card)
	game := blackjack.BlackJackGame{}
	game.GetPlayerInput(os.Stdin)
	fmt.Println("Game is: ", game)
}
