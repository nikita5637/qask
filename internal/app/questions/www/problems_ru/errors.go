package problems_ru

import (
	"errors"
)

var (
	ErrCouldNotGetProblemID = errors.New("Could not get math problem ID")
	ErrCouldNotGetProblem   = errors.New("Could not get math problem")
	ErrGetInvalidProblem    = errors.New("Get invalid problem")
)
