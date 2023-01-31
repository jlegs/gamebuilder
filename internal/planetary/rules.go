package planetary

import (
	"fmt"
	"os"
	"strconv"
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

type StealCard CardAbility

func (sc StealCard) Act(obj interface{}) {
	fmt.Println("INSIDE stealCard.Act")
	switch g := obj.(type) {
	case *PlanetaryGame:
		fmt.Println("Game was *planetaryGame")
		fmt.Println(g)
		// player := g.GetPlayerInput(os.Stdin)
		// fmt.Println("You'd like to target player: ", player)
		fmt.Println("Pick an index from 1 to ", len(g.Players[0].Hand.Cards)-1)
		index, err := strconv.Atoi(g.GetPlayerInput(os.Stdin))
		if err != nil {
			fmt.Println("\n------\nUH OH AN ERROR OH")
		}
		card := g.Players[0].Hand.Cards[index]
		fmt.Println("You picked: ", card)
		g.Players[0].Hand.RemoveCard(index)
		fmt.Println("Length of players cards is now: ", len(g.Players[0].Hand.Cards))

	}
}
