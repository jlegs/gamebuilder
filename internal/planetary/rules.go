package planetary

import (
	"fmt"
	"os"
)

type DrawFour CardAbility

func (df DrawFour) Name() string {
	return "DrawFour"
}

func (df DrawFour) Act(obj interface{}) {
	switch g := obj.(type) {
	case *PlanetaryGame:
		fmt.Println("Called drawfour.Act()!")
		c1 := g.PlanetaryDeck.DealCard()
		c2 := g.PlanetaryDeck.DealCard()
		c3 := g.PlanetaryDeck.DealCard()
		c4 := g.PlanetaryDeck.DealCard()
		g.Players[0].Hand.AddCard(c1)
		g.Players[0].Hand.AddCard(c2)
		g.Players[0].Hand.AddCard(c3)
		g.Players[0].Hand.AddCard(c4)
		fmt.Println("Nows the part where player input required: ")
		text := g.GetPlayerInput(os.Stdin)
		fmt.Println("RETRIEVED TEXT AND IT WAS: ", text)
	default:
		fmt.Println("Wrong type, not Planetary Game Object Ppointer.")
	}
}

func (df DrawFour) Description() {
	fmt.Println("Playing this card will allow you to draw four cards from the deck instantly.")
}
