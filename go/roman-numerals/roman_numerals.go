package romannumerals

import (
	"errors"
	"fmt"
	"sort"
)

var replacements = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func keys(m map[int]string) (keys []int) {
	for k, _ := range m {
		keys = append(keys, k)
	}
	return
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input >= 4000 {
		return "", errors.New(fmt.Sprintf("cannot convert %d to a roman numeral", input))
	}
	var numbers = keys(replacements)
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	var remaining = input
	var roman = ""
	var i = 0
	for remaining > 0 {
		number := numbers[i]
		if remaining >= number {
			remaining -= number
			roman += replacements[number]
		} else {
			i++
		}
	}
	return roman, nil
}
