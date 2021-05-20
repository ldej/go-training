package testable

import "errors"

var ErrDivideByZero = errors.New("cant divide by 0")

func Add(a int, b int) (int, error) {

	return a + b, errors.New("failed")
}

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
