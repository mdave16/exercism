package grains

import "errors"

// Square returns the number of grains on the nth square of a chess board.
func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("invalid tile on chess board")
	}
	return 1 << (number - 1), nil
}

// Total returns the total number of grains of rice on a chess board.
func Total() uint64 {
	return 1<<64 - 1
}
