package main

import (
	"fmt"
)

func processMessages(messages []string) []string {
	ch := make(chan string)
	processedMessages := make([]string, len(messages))
	go func() {
		for _, msg := range messages {
			fmt.Print(msg)
			ch <- msg + "-processed"
		}
	}()
	for i := 0; i < len(messages); i++ {
		processedMessages[i] = <-ch
	}
	close(ch)
	return processedMessages
}

func main() {
	messages := []string{"Here are some messages", "Here is another", "and another"}
	processedMessages := processMessages(messages)
	fmt.Print(processedMessages)
	// prints ["Here are some messages-processed", "Here is another-processed", "and another-processed"]

}
