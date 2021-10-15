// Package hello contains interfaces shared between plugins and the main program
package hello

// Hello is the interface implemented by each plugin
type Hello interface {
	Greeting() string
	Farewell() string
}
