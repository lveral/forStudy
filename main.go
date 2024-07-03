package main

import (
	"fmt"
	"testing"
)

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" ||
		mToSend.sender.number == 0 ||
		mToSend.recipient.name == "" ||
		mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func main() {
	var t *testing.T
	Test(t)
}

func Test(t *testing.T) {
	type testCase struct {
		mToSend  messageToSend
		expected bool
	}

	tests := []testCase{
		{messageToSend{"Thanks for signing up", user{"Name1", 148255510981}, user{"Name2", 148255510982}}, true},
		{messageToSend{"Thanks for signing up", user{"", 148255510981}, user{"Name2", 148255510982}}, false},
		{messageToSend{"Thanks for signing up", user{"Name1", 0}, user{"Name2", 148255510982}}, false},
		{messageToSend{"Thanks for signing up", user{"Name1", 148255510981}, user{"", 148255510982}}, false},
		{messageToSend{"Thanks for signing up", user{"Name1", 148255510981}, user{"Name2", 0}}, false},
		{messageToSend{"Thanks for signing up", user{"", 0}, user{"", 0}}, false},
	}

	for _, test := range tests {
		if output := canSendMessage(test.mToSend); output != test.expected {
			t.Errorf(
				"Test Failed: (%v, %v, %v) -> expected: %v actual: %v",
				test.mToSend.message,
				test.mToSend.sender,
				test.mToSend.recipient,
				test.expected,
				output,
			)
		} else {
			fmt.Printf("Test Passed: (%v, %v, %v) -> expected: %v actual: %v\n",
				test.mToSend.message,
				test.mToSend.sender,
				test.mToSend.recipient,
				test.expected,
				output,
			)
		}
	}
}
