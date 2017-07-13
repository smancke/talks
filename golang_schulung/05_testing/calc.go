package foo

import (
	"errors"
	"fmt"
)

type Operation func(a, b int) (int, error)

var operations map[string]Operation = map[string]Operation{
	"+": func(a, b int) (int, error) { return a + b, nil },
	"-": func(a, b int) (int, error) { return a - b, nil },
	"*": func(a, b int) (int, error) { return a * b, nil },
	"/": func(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("division by zero is undefined")
		}
		return a / b, nil
	},
}

func Calc(operation string, a, b int) (int, error) {
	op, found := operations[operation]
	if !found {
		return 0, fmt.Errorf("no such operation: %v", operation)
	}
	return op(a, b)
}
