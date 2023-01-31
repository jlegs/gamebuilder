package planetary

import (
	"bufio"
	"fmt"
	"gamebuilder/internal/cardgame"
	"io"
)

type PlanetaryGame struct {
	cardgame.CardGame
	CurrentTurn   cardgame.Player
	PlanetaryDeck *PlanetaryDeck
	MilitaryDeck  *MilitaryDeck
	Players       []PlanetaryPlayer
	Rulebook      PlanetaryRulebook
}

func (g *PlanetaryGame) Play() {

}

func (g *PlanetaryGame) Initialize(cardgame.Player) {

}

func (g *PlanetaryGame) EvaluateWinRules() string {

	return "unimplemented"
}

func (g *PlanetaryGame) ApplyRules() {

}

func (g *PlanetaryGame) Deal() {

}

func (g *PlanetaryGame) EndTurn() {

}

func (g *PlanetaryGame) StartTurn() {

}
func (g *PlanetaryGame) GetPlayerInput(stdin io.Reader) string {

	buf := bufio.NewScanner(stdin)
	for {
		buf.Scan()
		t := buf.Text()
		if t == "q" {
			break
		} else if t == "player" {
			fmt.Println("YOU ARE TARGETING A PLAYER!!!!!!")
			return t
		} else {
			return t
		}
	}
	return ""
}

type PlanetaryRulebook struct {
	WinningRules []cardgame.RuleCondition
}

type PlanetaryPlayer struct {
	cardgame.Player
	Hand  PlanetaryHand
	Cards []PlanetaryCard
}
