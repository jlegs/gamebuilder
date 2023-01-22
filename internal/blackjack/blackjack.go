package blackjack

import (
	"bufio"
	"gamebuilder/internal/cardgame"
	"io"
)

type BlackJackGame struct {
	// Deck cardgame.Deck
	cardgame.CardGame
	Rulebook *cardgame.Rulebook
	MinBet   int
	MaxBet   int
	Player   BJPlayer
	Dealer   *BJPlayer
	Deck     cardgame.Deck
	// SideGame cardgame.Game
}

func (bj *BlackJackGame) Initialize(player cardgame.Player) {
	// var Rules = BJRules
	rb := cardgame.Rulebook{
		Rules: []cardgame.RuleConditioner{HasBJ{}},
	}
	bj.Rulebook = &rb

	d := NewDeck(1)
	bj.Deck = d

	// for _, player := range players {
	bjplayer := BJPlayer{Name: player.GetName(), BankRoll: player.GetBankRoll()}
	newHand := BJHand{}
	bjplayer.Hand = newHand
	// }

	bj.Player = bjplayer

	bj.Dealer = &BJPlayer{Name: "Dealer"}

}

func (bj *BlackJackGame) Deal() {
	// c := Card{Value: "K", Suit: "h"}
	// allPlayers := []BJPlayer{}
	for i := 0; i < 2; i++ {
		// for _, player := range bj.Players {
		c := bj.Deck.DealCard()
		// fmt.Println(player)
		// fmt.Println(c)

		bj.Player.Hand.AddCard(&c)
		// }
	}
}

func (bj *BlackJackGame) Play() {

}

func (bj *BlackJackGame) ReceiveBets() {
	// should this be a channel?
	// would allow us to async receive bets ....
}

func (bj *BlackJackGame) GetPlayerInput(stdin io.Reader) string {

	buf := bufio.NewScanner(stdin)
	for {
		buf.Scan()
		t := buf.Text()
		if t == "q" {
			break
		} else {
			return t
		}
	}
	return ""
}

func (bj *BlackJackGame) EvaluateConditions() string {
	for _, rule := range bj.Rulebook.Rules {
		switch t := rule.(type) {
		case HasBJ:
			if t.Condition(bj) {
				return "IDK -- HasBJ was the rule condition that was gotten."
			}
		default:
			return "IDK"
		}
	}
	return "IDK"
}

func (bj *BlackJackGame) ApplyRules() {

}

type BJPlayer struct {
	// cardgame.Player
	Hand     BJHand
	Name     string
	BankRoll int
}

func (p *BJPlayer) GetName() string {
	return p.Name
}

func (p *BJPlayer) GetBankRoll() int {
	return p.BankRoll
}
