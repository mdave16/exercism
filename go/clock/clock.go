package clock

import "fmt"

type Clock struct {
	minute int
}

func postiveMod(b int) func(int) int {
	return func(a int) int {
		return ((a % b) + b) % b
	}
}

var inDay = postiveMod(24 * 60)

func New(h, m int) Clock {
	return Clock{0}.Add(inDay(h*60 + m))
}

func (c Clock) Add(m int) Clock {
	return Clock{inDay(c.minute + m)}
}

func (c Clock) Subtract(m int) Clock {
	return Clock{inDay(c.minute - m)}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minute/60, c.minute%60)
}
