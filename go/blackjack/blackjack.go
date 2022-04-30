package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	strToNum := 0

	switch card {
	case "ace":
		strToNum = 11
	case "king", "queen", "jack", "ten":
		strToNum = 10
	case "nine":
		strToNum = 9
	case "eight":
		strToNum = 8
	case "seven":
		strToNum = 7
	case "six":
		strToNum = 6
	case "five":
		strToNum = 5
	case "four":
		strToNum = 4
	case "three":
		strToNum = 3
	case "two":
		strToNum = 2
	default:
		return 0
	}
	return strToNum
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	playerAction := ""
	playerCardSum := ParseCard(card1) + ParseCard(card2)
	dealerCardSum := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace":
		playerAction = "P"
	case playerCardSum == 21 && dealerCardSum < 10:
		playerAction = "W"
	case (playerCardSum >= 17 && playerCardSum <= 21) || (playerCardSum >= 12 && playerCardSum <= 16 && dealerCardSum < 7):
		playerAction = "S"
	case (playerCardSum >= 12 && playerCardSum <= 16 && dealerCardSum >= 7) || playerCardSum <= 11:
		playerAction = "H"
	}
	return playerAction
}

/*
Stand (S)
Hit (H)
Split (P)
Automatically win (W)

If you have a pair of aces you must always split them.
If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
If your cards sum up to a value within the range [17, 20] you should always stand.
If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
If your cards sum up to 11 or lower you should always hit.
*/
