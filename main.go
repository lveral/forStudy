package main

import (
	"fmt"
)

type PlainText struct {
	message string
}

func (p PlainText) Format() string {
	return p.message
}

type Bold struct {
	message string
}

func (b Bold) Format() string {
	return fmt.Sprintf("**%s**", b.message)
}

type Code struct {
	message string
}

func (c Code) Format() string {
	return fmt.Sprintf("`%s`", c.message)
}

type Formatter interface {
	Format() string
}

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}

func main() {
	pt := SendMessage(PlainText{"M1"})
	fmt.Println(pt)
	b := SendMessage(Bold{"M1"})
	fmt.Println(b)
	c := SendMessage(Code{"M1"})
	fmt.Println(c)
}
