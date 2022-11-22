package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`([a-z0-9]+')?[a-z0-9]+`)
	freq := make(map[string]int)
	for _, word := range re.FindAll([]byte(strings.ToLower(phrase)), -1) {
		if len(word) == 0 {
			continue
		}
		freq[string(word)]++
	}
	return freq
}
