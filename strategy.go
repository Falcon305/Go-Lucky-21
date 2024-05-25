package main

func BasicStrategyAdvice(playerHand, dealerHand []Card) string {
	playerValue := HandValue(playerHand)
	dealerUpCard := dealerHand[0]

	switch {
	case playerValue >= 17:
		return "Stand"
	case playerValue >= 13 && CardValue(dealerUpCard) <= 6:
		return "Stand"
	case playerValue == 12 && CardValue(dealerUpCard) >= 4 && CardValue(dealerUpCard) <= 6:
		return "Stand"
	case playerValue == 11:
		return "Double Down if allowed, otherwise Hit"
	case playerValue == 10 && CardValue(dealerUpCard) <= 9:
		return "Double Down if allowed, otherwise Hit"
	case playerValue == 9 && CardValue(dealerUpCard) >= 3 && CardValue(dealerUpCard) <= 6:
		return "Double Down if allowed, otherwise Hit"
	case playerValue <= 8:
		return "Hit"
	default:
		return "Hit"
	}
}
