package main

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	var fm []Message
	for _, msg := range messages {
		if msg.Type() == filterType {
			fm = append(fm, msg)
		}
	}
	return fm
}

func main() {
	test := []Message{
		TextMessage{"S1", "content1"},
		MediaMessage{"S2", "type1", "content2"},
		LinkMessage{"S3", "URL1", "content3"},
		MediaMessage{"S4", "type2", "content4"},
		LinkMessage{"S5", "URL2", "content5"},
		LinkMessage{"S6", "URL3", "content6"},
	}
	fmt.Println(filterMessages(test, "text"))
	fmt.Println(filterMessages(test, "media"))
	fmt.Println(filterMessages(test, "link"))
}
