package main

import (
	"github.com/simonbrady/go-plugin-example/hello"
)

// Private type that implements the hello.Hello interface
type greeter struct{}

func (greeter) Greeting() string {
	return "hello from bar"
}

func (greeter) Farewell() string {
	return "goodbye from bar"
}

// Singleton instance of the type
var greeterInstance greeter

// GetHello returns our implementation of the hello.Hello interface
func GetHello() hello.Hello {
	return greeterInstance
}
