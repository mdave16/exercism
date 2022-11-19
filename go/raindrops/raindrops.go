package raindrops

import "fmt"

func Convert(number int) (sound string) {
	if number%3 == 0 {
		sound += "Pling"
	}
	if number%5 == 0 {
		sound += "Plang"
	}
	if number%7 == 0 {
		sound += "Plong"
	}
	if len(sound) == 0 {
		sound += fmt.Sprint(number)
	}
	return
}
