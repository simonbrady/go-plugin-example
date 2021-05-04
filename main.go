package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/simonbrady/go-plugin-example/hello"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	execPath, err := os.Executable()
	check(err)
	pluginPaths, err := filepath.Glob(filepath.Join(filepath.Dir(execPath), "plugins", "*.so"))
	check(err)
	var plugins []hello.Hello
	for _, pluginPath := range pluginPaths {
		fmt.Printf("Loading %s\n", pluginPath)
		p, err := plugin.Open(pluginPath)
		check(err)
		// Call factory method to get plugin's implementation of the hello.Hello interface
		getHello, err := p.Lookup("GetHello")
		check(err)
		plugins = append(plugins, getHello.(func() hello.Hello)())
	}
	// Invoke plugins through the common interface
	for _, p := range plugins {
		fmt.Println(p.Greeting())
		defer fmt.Println(p.Farewell())
	}
}
