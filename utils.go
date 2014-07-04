package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func cwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func relPath(path string) string {
	return "." + string(filepath.Separator) + path
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
