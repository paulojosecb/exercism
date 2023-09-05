package blackjack

var cards = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
	"other": 0}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := cards[card1] + cards[card2]
	dealeValue := cards[dealerCard]
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case sum == 21 && (dealerCard != "ace" && dealeValue != 10):
		return "W"
	case sum == 21 && dealeValue >= 10:
		return "S"
	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16 && dealeValue < 7:
		return "S"
	case sum >= 12 && sum <= 16 && dealeValue >= 7:
		return "H"
	case sum <= 11:
		return "H"
	default:
		return "S"
	}
}
