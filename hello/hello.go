package hello

// Hello is the interface implemented by each plugin
type Hello interface {
	Greeting() string
	Farewell() string
}
