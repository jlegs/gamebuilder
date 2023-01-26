package blackjack

import (
	"bufio"
	"fmt"
	"gamebuilder/internal/cardgame"
	"io"
	"os"
)

type BlackJackGame struct {
	// Deck cardgame.Deck
	cardgame.CardGame
	// Game     cardgame.Game
	Rulebook *cardgame.Rulebook
	MinBet   int
	MaxBet   int
	Players  []*BJPlayer
	Dealer   *BJPlayer
	Deck     cardgame.Deck
	// SideGame cardgame.Game
}

func (bj *BlackJackGame) Initialize(player cardgame.Player) {
	// var Rules = BJRules
	rb := cardgame.Rulebook{
		WinRules: []cardgame.RuleConditioner{HasBJ{}, Busted{}, Tie{}, PlayerStrongerHand{}},
	}
	bj.Rulebook = &rb

	d := NewDeck(1)
	bj.Deck = d

	// for _, player := range players {
	bjplayer := BJPlayer{Name: player.GetName(), BankRoll: player.GetBankRoll()}
	newHand := BJHand{}
	bjplayer.Hand = newHand
	// }

	bj.Players = append(bj.Players, &bjplayer)

	bj.Dealer = &BJPlayer{Name: "Dealer"}

}

func (bj *BlackJackGame) Deal() {
	// c := Card{Value: "K", Suit: "h"}
	allPlayers := append(bj.Players, bj.Dealer)
	for i := 0; i < 2; i++ {
		for _, player := range allPlayers {
			c := bj.Deck.DealCard()
			player.Hand.AddCard(&c)
		}
	}
}

func (bj *BlackJackGame) Play() {
	bj.Deck.Shuffle()
	bj.Deal()
	fmt.Println("YOU have: ")
	bj.Players[0].Hand.ShowHand()
	fmt.Println("Dealer has: ", bj.Dealer.Hand.Cards[0])
	for {
		if bj.Players[0].Hand.CalculateBJ() < 22 {
			bj.Players[0].Hand.ShowHand()
			fmt.Println("Enter 'h' to hit, or anything else to stay: \n>")
			input := bj.GetPlayerInput(os.Stdin)
			if input == "h" {
				card := bj.Deck.DealCard()
				bj.Players[0].Hand.AddCard(&card)
			} else if input == "j" {
				var playableCards []cardgame.Card
				for _, c := range bj.Players[0].Hand.Cards {
					card, _ := c.(*Card)
					for _, rule := range card.PlayRules {
						r, ok := rule.(*IsPlayable)
						if ok && r.Condition(c) {
							playableCards = append(playableCards, card)
							// card.Play()
							fmt.Println("End of the *Is10 switch case inside Card type assertion\n==========", &playableCards)
						}

					}
					// }
				}
			} else {
				break
			}
		} else {
			break
		}
	}

	for {
		if bj.Dealer.Hand.CalculateBJ() <= 16 {
			card := bj.Deck.DealCard()
			bj.Dealer.Hand.AddCard(&card)
		} else {
			break
		}
	}

	fmt.Println("Dealers cards were: ")
	bj.Dealer.Hand.ShowHand()
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

func (bj *BlackJackGame) EvaluateWinRules() string {
	switch t := bj.Players[0].Hand.Cards[0].(type) {
	case *Card:

		for _, rule := range t.PlayRules {
			if rule.Condition(t) {
				fmt.Println("CARD WAS A 10 CARD!!!")
			} else {
				fmt.Println("Card was not a 10 or logic is bad")
			}
		}
	}

	for _, rule := range bj.Rulebook.WinRules {
		switch t := rule.(type) {
		case HasBJ:
			if t.Condition(*bj) {
				return "PLAYER WON BECAUSE THEY HAD BLACKJACK"
			}
		case Busted:
			if t.Condition(*bj) {
				return "PLAYER LOST BECAUSE THEY BUSTED"
			}
		case Tie:
			if t.Condition(*bj) {
				return "PLAYER TIED BC EVEN CARDS"
			}
		case PlayerStrongerHand:
			if t.Condition(*bj) {
				return "PLAYER WON! YAY!"
			}
		default:
			return "IDK 1"
		}
	}
	return "IDK 2"
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
