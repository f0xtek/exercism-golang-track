// Package scrabble returns the scrabble score for a given word
package scrabble

import "unicode"

// Score iterates over the word to calculate & return the scrabble score
func Score(word string) int {
	totalScore := 0
	for _, i := range word {

		var score int

		switch unicode.ToUpper(i) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score = 1
		case 'D', 'G':
			score = 2
		case 'B', 'C', 'M', 'P':
			score = 3
		case 'F', 'H', 'V', 'W', 'Y':
			score = 4
		case 'K':
			score = 5
		case 'J', 'X':
			score = 8
		case 'Q', 'Z':
			score = 10
		}
		totalScore += score
	}
	return totalScore
}
