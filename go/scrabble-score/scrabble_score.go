package scrabble

import "strings"

var scoreToLetter = map[int][]rune{
	1:  []rune{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'},
	2:  []rune{'D', 'G'},
	3:  []rune{'B', 'C', 'M', 'P'},
	4:  []rune{'F', 'H', 'V', 'W', 'Y'},
	5:  []rune{'K'},
	8:  []rune{'J', 'X'},
	10: []rune{'Q', 'Z'},
}

// transpose will transpose the data of map[int][]rune to map[rune]int
func transpose(input map[int][]rune) map[rune]int {
	var out = map[rune]int{}
	for k, v := range input {
		for _, l := range v {
			out[l] = k
		}
	}
	return out
}

var letterToScore = transpose(scoreToLetter)

// Score calculates the scrabble score of a word.
func Score(word string) (sum int) {
	for _, l := range strings.ToUpper(word) {
		sum += letterToScore[l]
	}
	return
}
