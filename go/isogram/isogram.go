package isogram

import "strings"

func lettersOnly(word string) string {
	return strings.ToUpper(strings.ReplaceAll(
		strings.ReplaceAll(word, " ", ""), "-", ""))
}

func IsIsogram(word string) bool {
	var seen = map[rune]bool{}
	for _, r := range lettersOnly(word) {
		_, isSeen := seen[r]
		if isSeen {
			return false
		} else {
			seen[r] = true
		}
	}
	return true
}
