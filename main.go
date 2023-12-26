package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
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
	for _, pluginPath := range pluginPaths {
		fmt.Printf("Loading %s\n", pluginPath)
		p := must(plugin.Open(pluginPath))
		hello := must(p.Lookup("Hello"))
		fmt.Println(hello.(func() string)())
	}
}
