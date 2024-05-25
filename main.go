package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit string
type Value string

const (
	Spades   Suit = "Spades"
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
)

const (
	Ace   Value = "Ace"
	Two   Value = "2"
	Three Value = "3"
	Four  Value = "4"
	Five  Value = "5"
	Six   Value = "6"
	Seven Value = "7"
	Eight Value = "8"
	Nine  Value = "9"
	Ten   Value = "10"
	Jack  Value = "Jack"
	Queen Value = "Queen"
	King  Value = "King"
)

type Card struct {
	Suit  Suit
	Value Value
}

type Deck []Card

func NewDeck() Deck {
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	values := []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	var deck Deck
	for _, suit := range suits {
		for _, value := range values {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}
	return deck
}

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

func (d *Deck) Deal() Card {
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}

func CardValue(card Card) int {
	switch card.Value {
	case Ace:
		return 11
	case Two:
		return 2
	case Three:
		return 3
	case Four:
		return 4
	case Five:
		return 5
	case Six:
		return 6
	case Seven:
		return 7
	case Eight:
		return 8
	case Nine:
		return 9
	case Ten, Jack, Queen, King:
		return 10
	}
	return 0
}

func HandValue(hand []Card) int {
	value := 0
	aces := 0
	for _, card := range hand {
		if card.Value == Ace {
			aces++
		}
		value += CardValue(card)
	}
	for value > 21 && aces > 0 {
		value -= 10
		aces--
	}
	return value
}

var runningCount int

func UpdateCount(card Card) {
	switch card.Value {
	case Two, Three, Four, Five, Six:
		runningCount++
	case Ten, Jack, Queen, King, Ace:
		runningCount--
	}
}

func main() {
	deck := NewDeck()
	deck.Shuffle()
	runningCount = 0

	wins, losses, ties := 0, 0, 0

	for {
		playerHand := []Card{deck.Deal(), deck.Deal()}
		dealerHand := []Card{deck.Deal(), deck.Deal()}

		// Update running count
		for _, card := range playerHand {
			UpdateCount(card)
		}
		for _, card := range dealerHand {
			UpdateCount(card)
		}

		fmt.Println("Player Hand:", playerHand)
		fmt.Println("Dealer Hand:", dealerHand[0], "Hidden")

		playerTurn := true
		for playerTurn {
			fmt.Println("Player Hand Value:", HandValue(playerHand))
			var action string
			fmt.Print("Do you want to (h)it or (s)tand? ")
			fmt.Scanf("%s", &action)
			if action == "h" {
				card := deck.Deal()
				playerHand = append(playerHand, card)
				UpdateCount(card)
				if HandValue(playerHand) > 21 {
					fmt.Println("Player busts! Dealer wins.")
					losses++
					playerTurn = false
				}
			} else {
				playerTurn = false
			}
		}

		fmt.Println("Dealer reveals second card:", dealerHand[1])
		for HandValue(dealerHand) < 17 {
			card := deck.Deal()
			dealerHand = append(dealerHand, card)
			UpdateCount(card)
		}
		playerValue := HandValue(playerHand)
		dealerValue := HandValue(dealerHand)
		fmt.Println("Dealer Hand Value:", dealerValue)
		if playerValue <= 21 && (dealerValue > 21 || playerValue > dealerValue) {
			fmt.Println("Player wins!")
			wins++
		} else if playerValue == dealerValue {
			fmt.Println("Push!")
			ties++
		} else {
			fmt.Println("Dealer wins!")
			losses++
		}

		fmt.Printf("Running Count: %d\n", runningCount)
		fmt.Printf("Wins: %d, Losses: %d, Ties: %d\n", wins, losses, ties)
		var again string
		fmt.Print("Play again? (y/n): ")
		fmt.Scanf("%s", &again)
		if again != "y" {
			break
		}
	}
}
