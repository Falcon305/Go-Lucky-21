package main

import (
	"fmt"
)

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

func UpdateCount(runningCount *int, card Card) {
	switch card.Value {
	case Two, Three, Four, Five, Six:
		*runningCount++
	case Ten, Jack, Queen, King, Ace:
		*runningCount--
	}
}

func PrintHand(label string, hand []Card) {
	fmt.Printf("%s%s:%s ", Cyan, label, Reset)
	for _, card := range hand {
		fmt.Printf("%s%s of %s%s, ", Magenta, card.Value, card.Suit, Reset)
	}
	fmt.Println()
}
