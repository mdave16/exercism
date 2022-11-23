package isbn

import "strconv"

func IsValidISBN(stringIsbn string) bool {
	isbn := []int{}
	for i, d := range stringIsbn {
		if d == '-' {
			continue
		}
		if d == 'X' && i == (len(stringIsbn)-1) {
			isbn = append(isbn, 10)
			continue
		}
		n, err := strconv.Atoi(string(d))
		if err != nil {
			return false
		}
		isbn = append(isbn, n)
	}
	if len(isbn) < 9 {
		return false
	}
	return CheckValue(isbn[:len(isbn)-1]) == isbn[len(isbn)-1]
}

type checkFn func(sum, d, i int) int

func isbn10check(sum, d, i int) int {
	sum += (d * (i + 1))
	sum %= 11
	return sum
}
func isbn13check(sum, d, i int) int {
	sum += (d * (1 + 2*(i%2)))
	sum %= 10
	return sum
}
func CheckValue(incompleteIsbn []int) int {
	var check checkFn
	if len(incompleteIsbn) == 9 {
		check = isbn10check
	} else if len(incompleteIsbn) == 12 {
		check = isbn13check
	} else {
		return -1
	}
	sum := 0
	for i, d := range incompleteIsbn {
		sum = check(sum, d, i)
	}
	return sum
}
