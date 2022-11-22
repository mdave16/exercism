package anagram

import (
	"sort"
	"strings"
)

type ByRune []rune

func (a ByRune) Len() int           { return len(a) }
func (a ByRune) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRune) Less(i, j int) bool { return a[i] < a[j] }

func IsAnagram(subject, candidate string) bool {
	if len(subject) != len(candidate) {
		return false
	}
	subj := strings.ToUpper(subject)
	cand := strings.ToUpper(candidate)
	if subj == cand {
		return false
	}
	subjectRunes:= []rune(subj)
	candidateRunes:=[]rune(cand)
	sort.Sort(ByRune(subjectRunes))
	sort.Sort(ByRune(candidateRunes))
	return string(subjectRunes) == string(candidateRunes)
}

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	for _, candidate := range candidates {
		if IsAnagram(subject, candidate) {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}
