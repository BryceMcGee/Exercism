package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.

func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	switch {
	case err == ErrScaleMalfunction && fodder > 0:
		fodder *= 2
	case err != nil && err != ErrScaleMalfunction:
		return 0, err
	case fodder < 0:
		return 0, errors.New("negative fodder")
	case cows == 0:
		return 0, errors.New("division by zero")
	case cows < 0:
		return 0, &SillyNephewError{cows: cows}
	}
	return fodder / float64(cows), nil
}
