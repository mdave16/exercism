package luhn

import "strconv"

func Valid(id string) bool {

	sum := 0
	shouldDouble := false
	validLength := false

	for i := 0; i < len(id); i++ {
		digit := rune(id[len(id)-1-i])
		if digit == ' ' {
			continue
		}
		n, err := strconv.Atoi(string(digit))
		if err != nil {
			return false
		}
		if shouldDouble {
			if !validLength {
				validLength = true
			}
			n = n * 2
			if n > 9 {
				n = n - 9
			}
		}
		shouldDouble = !shouldDouble
		sum = sum + n

	}

	return validLength && (sum%10 == 0)
}
