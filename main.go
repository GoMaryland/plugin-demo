package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	fmt.Println(getPlugins())
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
