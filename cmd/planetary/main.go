package main

import (
	"fmt"
	"gamebuilder/internal/planetary"
)

func main() {
	deck := planetary.NewDeck(1)
	fmt.Println("Instantiating Deck")
	fmt.Println(deck)

	card := deck.Cards[0]
	card.Ability.Act(deck)
	// switch ability := card.Ability.(type) {
	// case planetary.DrawFour:
	// 	fmt.Println("Case was DrawFour")
	// 	ability.Act(deck)
	// }
}
