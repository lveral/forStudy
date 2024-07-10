package main

import (
	"fmt"
	"errors"
)

type divideError struct {
	dividend float64
}

func (d divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}
	

func validateStatus(status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	}
	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}

func main() {
	tests := []string{
		"",
		"normal status",
		"very big status very big status very big status very big status very big status very big status very big status very big status very big status ",
	}
	for _, test := range tests {
		fmt.Println(validateStatus(test))
	}
  ///other test
  a := 10.0
	b := 0.0
	fmt.Println(divide(a, b))
}
