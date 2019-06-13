package main

import (
	"github.com/simonbrady/go-plugin-example/hello"
)

// Private type that implements the hello.Hello interface
type greeter struct {}

func (greeter) Greeting() string {
	return "hello from foo"
}

func (greeter) Farewell() string {
	return "goodbye from foo"
}

// Singleton instance of the type
var greeter_instance greeter

// Factory method returning our implementation of the interface
func GetHello() hello.Hello {
	return greeter_instance
}
