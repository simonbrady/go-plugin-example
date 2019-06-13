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
	exec_path, err := os.Executable()
	check(err)
	plugin_dir := filepath.Join(filepath.Dir(exec_path), "plugins")
	plugin_paths, err := filepath.Glob(filepath.Join(plugin_dir, "*.so"))
	check(err)
	for _, plugin_path := range plugin_paths {
		fmt.Printf("Loading %s\n", plugin_path)
		p, err := plugin.Open(plugin_path)
		check(err)
		hello, err := p.Lookup("Hello")
		check(err)
		fmt.Println(hello.(func() string)())
	}
}
