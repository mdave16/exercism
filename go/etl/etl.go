package etl

import "strings"

// Transform will transform a scrabble score map
func Transform(scoreToLetters map[int][]string) map[string]int {
	var letterToScore = make(map[string]int)
	for score, letters := range scoreToLetters {
		for _, letter := range letters {
			letterToScore[strings.ToLower(letter)] = score
		}
	}
	return letterToScore
}
