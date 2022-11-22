package reverse

func Reverse(input string) string {
	runes := []rune(input)
	numRunes := len(runes)
	output := make([]rune, numRunes)
	for i, l := range runes {
		output[numRunes-i-1] = l
	}
	return string(output)
}
