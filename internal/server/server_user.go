package server

import "fmt"

type Message string

func NewMessage(message string) Message {
	return Message(fmt.Sprintf("Hi, %v", message))
}

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}
