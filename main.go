package main

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
	Membership
}

func (user User) SendMessage(message string, messageLength int) (string, bool) {
	if user.MessageCharLimit >= messageLength {
		return message, true
	}
	return "", false
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

func newUser(name string, membershipType membershipType) User {
	membership := Membership{Type: membershipType}
	if membershipType == TypeStandard {
		membership.MessageCharLimit = 100
	} else if membershipType == TypePremium {
		membership.MessageCharLimit = 1000
	}
	return User{Name: name, Membership: membership}
}

func main() {
	var t *testing.T
	Test(t)
}

func Test(t *testing.T) {
	type expectedCase struct {
		expectedMessage string
		result          bool
	}
	type testCase struct {
		message       string
		messageLength int
		user          User
		expectedCase
	}

	tests := []testCase{
		{"message1", 100, User{Name: "Name1", Membership: Membership{TypeStandard, 100}}, expectedCase{"message1", true}},
		{"message2", 110, User{Name: "Name1", Membership: Membership{TypeStandard, 100}}, expectedCase{"", false}},
		{"message3", 1000, User{Name: "Name1", Membership: Membership{TypePremium, 1000}}, expectedCase{"message3", true}},
		{"message4", 1100, User{Name: "Name1", Membership: Membership{TypePremium, 1000}}, expectedCase{"", false}},
	}

	for _, test := range tests {
		if output, res := test.user.SendMessage(test.message, test.messageLength); output != test.expectedMessage || res != test.result {
			t.Errorf(
				"Test Failed: (%v, %v, %v) -> expected: %v, %v actual: %v, %v",
				test.message,
				test.messageLength,
				test.user,
				test.expectedMessage,
				test.result,
				output,
				res,
			)
		} else {
			fmt.Printf("Test Passed: (%v, %v, %v) -> expected: %v, %v actual: %v, %v\n",
				test.message,
				test.messageLength,
				test.user,
				test.expectedMessage,
				test.result,
				output,
				res,
			)
		}
	}
}
