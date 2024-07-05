package main

import (
	"errors"
	"fmt"
)

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
}
