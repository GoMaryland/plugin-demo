package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"regexp"
)

func main() {
	pluginPaths := getPlugins()
	var plugins []*plugin.Plugin
	for _, path := range pluginPaths {
		p, err := plugin.Open("plugins/" + path)
		if err != nil {
			panic(err)
		}
		plugins = append(plugins, p)
	}

	addFunc := findPlugin(plugins, "add")
	subtractFunc := findPlugin(plugins, "subtract")

	fmt.Println(addFunc(1, 1))
	fmt.Println(subtractFunc(2, 1))
}

func findPlugin(plugins []*plugin.Plugin, name string) func(int, int) int {
	for _, plugin := range plugins {
		opSymbol, err := plugin.Lookup("OpName")
		if err != nil {
			panic(err)
		}
		op := opSymbol.(*string)

		if *op == name {
			funcSymbol, err := plugin.Lookup("Operation")
			if err != nil {
				panic(err)
			}
			f := funcSymbol.(func(int, int) int)
			return f
		}
	}
	return nil
}

func getPlugins() []string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	filepath.Walk(cwd, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(".so", f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}
