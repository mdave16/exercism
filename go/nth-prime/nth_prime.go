package prime

import "errors"

var primes []int = []int{2}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("there is no zeroth/negative prime")
	}
	for len(primes) < n {
		lastPrime := primes[len(primes)-1]
		for i := (lastPrime + 1); i < lastPrime*2; i++ {
			divisible := false
			for _, p := range primes {
				if i%p == 0 {
					divisible = true
					continue
				}
			}
			if divisible {
				continue
			}
			primes = append(primes, i)
		}
	}
	return primes[n-1], nil
}
