package proverb

import "fmt"

func Proverb(rhyme []string) (proverb []string) {
	var line string
	for i := range rhyme {
		if i+1 < len(rhyme) {
			line = fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1])
		} else {
			line = fmt.Sprintf("And all for the want of a %v.", rhyme[0])
		}
		proverb = append(proverb, line)
	}
	return
}
