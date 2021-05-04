package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
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
	for _, pluginPath := range pluginPaths {
		fmt.Printf("Loading %s\n", pluginPath)
		p, err := plugin.Open(pluginPath)
		check(err)
		hello, err := p.Lookup("Hello")
		check(err)
		fmt.Println(hello.(func() string)())
	}
}
