package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

/*--------------------------------------
VARS
---------------------------------------*/
var fileTypesRegex string = ".(png|css|jpg|JPG|gif|GIF|doc|DOC|pdf|PDF|mov|mp4|mpg|wmv|webm|psd|m4v|jar|3gp|swf|SWF|eot|svg|ttf|woff|mp3)$"
var reporter = Reporter{removed: 0, remaining: 0, initial: 0}

/*--------------------------------------
types
---------------------------------------*/
type WalkerCb func(path string, bytes int64, match bool)

/*--------------------------------------
main
---------------------------------------*/
func main() {
	// dir := os.TempDir() + "filecleaner"
	dir := cwd()
	fmt.Printf("\nCleaning dir: (%s)\n\n", dir)
	Clean(dir)
	reporter.report()
}

/*--------------------------------------
file walker
---------------------------------------*/
// func vistorGenerator(cb func(path string, bytes int64, match bool)) func(path string, info os.FileInfo, err error) error {
func vistorGenerator(cb WalkerCb) func(path string, info os.FileInfo, err error) error {

	return func(path string, info os.FileInfo, err error) error {

		// only check files
		if !info.IsDir() {

			file, err := os.Open(path)
			fileStat, err := file.Stat()
			match, err := regexp.MatchString(fileTypesRegex, path)

			if err != nil {
				fmt.Println("error")
				fmt.Println(err)
				return err
			} else {
				// fmt.Println(path, fileStat.Size(), match)
				cb(path, fileStat.Size(), match)
			}
		}
		return nil
	}
}

func cleanWalkCb(path string, bytes int64, match bool) {
	reporter.increment(path, bytes, match)
	if match {
		os.Remove(path)
	}
}

func Clean(dir string) {
	err := filepath.Walk(dir, vistorGenerator(cleanWalkCb))
	if err != nil {
		fmt.Println(err)
	}
}
