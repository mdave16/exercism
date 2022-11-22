package pangram

import "strings"

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func contains(haystack string, needle rune)bool{
	for _, l := range haystack {
		if l == needle {
			return true
		}
	}
	return false
}

func IsPangram(input string) bool {
	pangram := strings.ToLower(input)
	for _,l := range alphabet {
		if(!contains(pangram, l)) {
			return false
		}
	}
	return true
}
