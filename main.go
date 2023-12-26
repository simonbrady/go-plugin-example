package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/simonbrady/go-plugin-example/hello"
)

// Generic wrapper for any function with error as its second return type,
// so the caller doesn't have to worry about error handling.
func must[T any](retval T, err error) T {
	if err != nil {
		panic(err)
	}
	return retval
}

func main() {
	execPath := must(os.Executable())
	pluginPaths := must(filepath.Glob(filepath.Join(filepath.Dir(execPath), "plugins", "*.so")))
	var plugins []hello.Hello
	for _, pluginPath := range pluginPaths {
		fmt.Printf("Loading %s\n", pluginPath)
		p := must(plugin.Open(pluginPath))
		// Call factory method to get plugin's implementation of the hello.Hello interface
		getHello := must(p.Lookup("GetHello"))
		plugins = append(plugins, getHello.(func() hello.Hello)())
	}
	// Invoke plugins through the common interface
	for _, p := range plugins {
		fmt.Println(p.Greeting())
		defer fmt.Println(p.Farewell())
	}
}
