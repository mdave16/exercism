package triangle

import "math"

type Kind int

const (
	NaT = iota
	Equ
	Iso
	Sca
)

// KindFromSides returns what kind of triangle is produced by having three sides of a given length.
func KindFromSides(a, b, c float64) Kind {
	max := math.Max(math.Max(a, b), c)
	min := math.Min(math.Min(a, b), c)
	mid := a + b + c - max - min
	if min+mid <= max {
		return NaT
	}
	if max == min {
		return Equ
	}
	if max == mid || mid == min {
		return Iso
	}
	return Sca
}
