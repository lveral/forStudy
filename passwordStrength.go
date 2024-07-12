package main

import (
	"fmt"
	"unicode"
)

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	var hasDigit, hasUpper bool
	for _, r := range password {
		if unicode.IsDigit(r) {
			hasDigit = true
		}
		if unicode.IsUpper(r) {
			hasUpper = true
		}
	}
	if !hasDigit || !hasUpper {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValidPassword("fail"))
	fmt.Println(isValidPassword("Ffail"))
	fmt.Println(isValidPassword("fail3"))
	fmt.Println(isValidPassword("failfailfailfail"))
	fmt.Println(isValidPassword("Pass1"))
}
