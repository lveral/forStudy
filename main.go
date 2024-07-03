package main

import (
	"fmt"
	"testing"
)

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType membershipType) User {
	var limit int
	if membershipType == TypeStandard {
		limit = 100
	} else {
		limit = 1000
	}
	return User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: limit,
		},
	}
}

func main() {
	var t *testing.T
	Test(t)
}

func Test(t *testing.T) {
	type testCase struct {
		name           string
		membershipType membershipType
		expected       User
	}

	tests := []testCase{
		{"Name1", TypeStandard, User{Name: "Name1", Membership: Membership{TypeStandard, 100}}},
		{"Name1", TypePremium, User{Name: "Name1", Membership: Membership{TypePremium, 1000}}},
	}

	for _, test := range tests {
		if output := newUser(test.name, test.membershipType); output != test.expected {
			t.Errorf(
				"Test Failed: (%v, %v) -> expected: %v actual: %v",
				test.name,
				test.membershipType,
				test.expected,
				output,
			)
		} else {
			fmt.Printf("Test Passed: (%v, %v) -> expected: %v actual: %v\n",
				test.name,
				test.membershipType,
				test.expected,
				output,
			)
		}
	}
}
