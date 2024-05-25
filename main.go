package main

import (
	"fmt"
)

func main() {
	stats, err := LoadStatistics()
	if err != nil {
		fmt.Printf("%sError loading statistics: %s%s\n", Red, err, Reset)
	}

	var mode string
	ClearScreen()
	fmt.Println("Select Mode:")
	fmt.Println("1. Game Mode")
	fmt.Println("2. Practice Mode")
	fmt.Scanln(&mode)

	for {
		deck := NewDeck()
		deck.Shuffle()
		runningCount := 0

		playerHand := []Card{deck.Deal(), deck.Deal()}
		dealerHand := []Card{deck.Deal(), deck.Deal()}

		for _, card := range playerHand {
			UpdateCount(&runningCount, card)
		}
		for _, card := range dealerHand {
			UpdateCount(&runningCount, card)
		}

		ClearScreen()
		PrintHand("Player Hand", playerHand)
		fmt.Printf("Dealer Hand: %s%s%s Hidden\n", Magenta, dealerHand[0], Reset)

		playerTurn := true
		for playerTurn {
			fmt.Printf("Player Hand Value: %s%d%s\n", Yellow, HandValue(playerHand), Reset)
			if mode == "2" { // Practice mode
				advice := BasicStrategyAdvice(playerHand, dealerHand)
				fmt.Printf("%sBasic Strategy Advice: %s%s\n", Blue, advice, Reset)
			}
			var action string
			fmt.Print("Do you want to (h)it or (s)tand? ")
			fmt.Scanln(&action)
			if action == "h" {
				card := deck.Deal()
				playerHand = append(playerHand, card)
				UpdateCount(&runningCount, card)
				if HandValue(playerHand) > 21 {
					fmt.Printf("%sPlayer busts! Dealer wins.%s\n", Red, Reset)
					stats.Losses++
					stats.MoneyLost += 10 // Assume each game has a bet of $10
					playerTurn = false
				}
			} else {
				playerTurn = false
			}
		}

		if HandValue(playerHand) <= 21 {
			ClearScreen()
			PrintHand("Player Hand", playerHand)
			PrintHand("Dealer Hand", dealerHand)

			for HandValue(dealerHand) < 17 {
				card := deck.Deal()
				dealerHand = append(dealerHand, card)
				UpdateCount(&runningCount, card)
			}

			playerValue := HandValue(playerHand)
			dealerValue := HandValue(dealerHand)
			fmt.Printf("Dealer Hand Value: %s%d%s\n", Yellow, dealerValue, Reset)

			if dealerValue > 21 || playerValue > dealerValue {
				fmt.Printf("%sPlayer wins!%s\n", Green, Reset)
				stats.Wins++
			} else if playerValue == dealerValue {
				fmt.Printf("%sPush!%s\n", Blue, Reset)
				stats.Ties++
			} else {
				fmt.Printf("%sDealer wins!%s\n", Red, Reset)
				stats.Losses++
				stats.MoneyLost += 10 // Assume each game has a bet of $10
			}
		}

		if mode == "2" { // Practice mode
			var userCount int
			fmt.Print("Enter your running count: ")
			fmt.Scanln(&userCount)
			if userCount == runningCount {
				fmt.Printf("%sCorrect count!%s\n", Green, Reset)
				stats.CorrectCounts++
				stats.CardCountStreak++
				if stats.CardCountStreak > stats.MaxCardCountStreak {
					stats.MaxCardCountStreak = stats.CardCountStreak
				}
			} else {
				fmt.Printf("%sIncorrect count. Correct count is %d.%s\n", Red, runningCount, Reset)
				stats.IncorrectCounts++
				stats.CardCountStreak = 0
			}
		}

		fmt.Printf("Running Count: %s%d%s\n", Yellow, runningCount, Reset)
		fmt.Printf("Wins: %s%d%s, Losses: %s%d%s, Ties: %s%d%s, Money Lost: %s$%d%s\n",
			Green, stats.Wins, Reset, Red, stats.Losses, Reset, Blue, stats.Ties, Reset, Red, stats.MoneyLost, Reset)
		fmt.Printf("Hands Played: %s%d%s, Correct Counts: %s%d%s, Incorrect Counts: %s%d%s\n",
			Yellow, stats.HandsPlayed, Reset, Green, stats.CorrectCounts, Reset, Red, stats.IncorrectCounts, Reset)
		fmt.Printf("Current Streak: %s%d%s, Max Streak: %s%d%s\n",
			Cyan, stats.CardCountStreak, Reset, Magenta, stats.MaxCardCountStreak, Reset)

		var again string
		fmt.Print("Play again? (y/n): ")
		fmt.Scanln(&again)
		if again != "y" {
			break
		}
		stats.HandsPlayed++
	}

	err = SaveStatistics(stats)
	if err != nil {
		fmt.Printf("%sError saving statistics: %s%s\n", Red, err, Reset)
	}
}
