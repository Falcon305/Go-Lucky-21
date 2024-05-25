package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
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

var (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
)

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

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printHand(label string, hand []Card) {
	fmt.Printf("%s%s:%s ", cyan, label, reset)
	for _, card := range hand {
		fmt.Printf("%s%s of %s%s, ", magenta, card.Value, card.Suit, reset)
	}
	fmt.Println()
}

func main() {
	var mode string
	clearScreen()
	fmt.Println("Select Mode:")
	fmt.Println("1. Game Mode")
	fmt.Println("2. Practice Mode")
	fmt.Scanln(&mode)

	moneyLost := 0
	wins, losses, ties := 0, 0, 0

	for {
		deck := NewDeck()
		deck.Shuffle()
		runningCount = 0

		playerHand := []Card{deck.Deal(), deck.Deal()}
		dealerHand := []Card{deck.Deal(), deck.Deal()}

		for _, card := range playerHand {
			UpdateCount(card)
		}
		for _, card := range dealerHand {
			UpdateCount(card)
		}

		clearScreen()
		printHand("Player Hand", playerHand)
		fmt.Printf("Dealer Hand: %s%s%s Hidden\n", magenta, dealerHand[0], reset)

		playerTurn := true
		for playerTurn {
			fmt.Printf("Player Hand Value: %s%d%s\n", yellow, HandValue(playerHand), reset)
			var action string
			fmt.Print("Do you want to (h)it or (s)tand? ")
			fmt.Scanln(&action)
			if action == "h" {
				card := deck.Deal()
				playerHand = append(playerHand, card)
				UpdateCount(card)
				if HandValue(playerHand) > 21 {
					fmt.Printf("%sPlayer busts! Dealer wins.%s\n", red, reset)
					losses++
					moneyLost += 10
					playerTurn = false
				}
			} else {
				playerTurn = false
			}
		}

		if HandValue(playerHand) <= 21 {
			clearScreen()
			printHand("Player Hand", playerHand)
			printHand("Dealer Hand", dealerHand)

			for HandValue(dealerHand) < 17 {
				card := deck.Deal()
				dealerHand = append(dealerHand, card)
				UpdateCount(card)
			}

			playerValue := HandValue(playerHand)
			dealerValue := HandValue(dealerHand)
			fmt.Printf("Dealer Hand Value: %s%d%s\n", yellow, dealerValue, reset)

			if dealerValue > 21 || playerValue > dealerValue {
				fmt.Printf("%sPlayer wins!%s\n", green, reset)
				wins++
			} else if playerValue == dealerValue {
				fmt.Printf("%sPush!%s\n", blue, reset)
				ties++
			} else {
				fmt.Printf("%sDealer wins!%s\n", red, reset)
				losses++
				moneyLost += 10
			}
		}

		if mode == "2" {
			var userCount int
			fmt.Print("Enter your running count: ")
			fmt.Scanln(&userCount)
			if userCount == runningCount {
				fmt.Printf("%sCorrect count!%s\n", green, reset)
			} else {
				fmt.Printf("%sIncorrect count. Correct count is %d.%s\n", red, runningCount, reset)
			}
		}

		fmt.Printf("Running Count: %s%d%s\n", yellow, runningCount, reset)
		fmt.Printf("Wins: %s%d%s, Losses: %s%d%s, Ties: %s%d%s, Money Lost: %s$%d%s\n",
			green, wins, reset, red, losses, reset, blue, ties, reset, red, moneyLost, reset)

		var again string
		fmt.Print("Play again? (y/n): ")
		fmt.Scanln(&again)
		if again != "y" {
			break
		}
	}
}
