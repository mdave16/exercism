package thefarm

// See types.go for the types defined for this exercise.
import (
	"errors"
	"fmt"
)

// SillyNephewError for when my silly nephew says there are negative cows.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, fodderError := weightFodder.FodderAmount()
	if fodderError != nil && fodderError != ErrScaleMalfunction {
		return 0, fodderError
	}
	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}
	if fodderError == ErrScaleMalfunction {
		fodder *= 2
	}
	return fodder / float64(cows), nil
}
