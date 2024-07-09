package main

import (
	"fmt"
)

type SomeInterface interface {
	FirstFunc() string
	//SecondFunc() string
}

type divideError struct {
	dividend float64
}

func (d divideError) FirstFunc() string {
	return fmt.Sprintf("FirstFunc %v", d.dividend)
}

func (d divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", d.dividend)
}

func (d divideError) SecondFunc() string {
	return fmt.Sprintf("SecondFunc %v", d.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func main() {
	a := 10.0
	b := 0.0
	fmt.Println(divide(a, b))
}
