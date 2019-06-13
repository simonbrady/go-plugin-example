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
	exec_path, err := os.Executable()
	check(err)
	plugin_paths, err := filepath.Glob(filepath.Join(filepath.Dir(exec_path), "plugins", "*.so"))
	check(err)
	var plugins []hello.Hello
	for _, plugin_path := range plugin_paths {
		fmt.Printf("Loading %s\n", plugin_path)
		p, err := plugin.Open(plugin_path)
		check(err)
		// Call factory method to get plugin's implementation of the hello.Hello interface
		get_hello, err := p.Lookup("GetHello")
		check(err)
		plugins = append(plugins, get_hello.(func() hello.Hello)())
	}
	// Invoke plugins through the common interface
	for _, p := range plugins {
		fmt.Println(p.Greeting())
		defer fmt.Println(p.Farewell())
	}
}
