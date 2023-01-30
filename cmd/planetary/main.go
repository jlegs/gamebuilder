package main

import (
	"fmt"
	"gamebuilder/internal/planetary"
)

func main() {

	fmt.Println("Instantiating Game")
	newHand := planetary.PlanetaryHand{}
	player := planetary.PlanetaryPlayer{Hand: newHand}
	g := planetary.PlanetaryGame{Players: []planetary.PlanetaryPlayer{player}}
	deck := planetary.NewDeck(1)
	g.PlanetaryDeck = deck
	g.PlanetaryDeck.Shuffle()
	milDeck := planetary.NewMilitaryDeck(1)
	g.MilitaryDeck = milDeck
	g.MilitaryDeck.Shuffle()

	g.PlanetaryDeck.PlanetaryCards[0].Ability.Act(&g)

	g.Players[0].Hand.AddCard(g.MilitaryDeck.DealCard())
	g.Players[0].Hand.ShowHand()
}
