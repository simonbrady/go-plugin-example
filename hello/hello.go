package hello

// Interface implemented by each plugin
type Hello interface {
	Greeting() string
	Farewell() string
}
