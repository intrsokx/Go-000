//package main
//
//import (
//	"github.com/intrsokx/Go-000/Week04/internal"
//	"log"
//)
//
//func main() {
//	log.Fatal(internal.NewApp().Run())
//}

package main

import (
	"fmt"
	"github.com/intrsokx/Go-000/Week04/internal"
	"log"
)

type Message struct {
	msg string
}
type Greeter struct {
	Message Message
}
type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
func (g Greeter) Greet() Message {
	return g.Message
}

// 使用wire前
func main() {
	log.Fatal(internal.NewApp().Run())
}

/*
// 使用wire后
func main() {
	event := InitializeEvent("hello_world")

	event.Start()
}*/
