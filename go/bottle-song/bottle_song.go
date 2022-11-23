package bottlesong

import "fmt"

func Recite(startBottles, takeDown int) []string {
	song := []string{}
	for bottles := startBottles; bottles > startBottles-takeDown; bottles-- {
		if bottles < startBottles {
			song = append(song, "")
		}
		song = append(song, []string{
			fmt.Sprintf("%s green %s hanging on the wall,", Title(num(bottles)), plural(bottles)),
			fmt.Sprintf("%s green %s hanging on the wall,", Title(num(bottles)), plural(bottles)),
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green %s hanging on the wall.", num(bottles-1), plural(bottles-1)),
		}...)
	}
	return song
}

func plural(num int) string {
	if num == 1 {
		return "bottle"
	}
	return "bottles"
}

func num(num int) string {
	return []string{
		"no",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	}[num]
}
