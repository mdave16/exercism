package raindrops

import "fmt"

var rules = []struct {
	divides int
	sound   string
}{{
	divides: 3,
	sound:   "Pling",
},
	{
		divides: 5,
		sound:   "Plang",
	},
	{
		divides: 7,
		sound:   "Plong",
	},
}

func Convert(number int) (sound string) {
	for _, rule := range rules {
		if number%rule.divides == 0 {
			sound += rule.sound
		}
	}
	if len(sound) == 0 {
		sound += fmt.Sprint(number)
	}
	return
}
