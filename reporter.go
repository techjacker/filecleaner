package main

import (
	"fmt"
)

type Reporter struct {
	removed, remaining, initial int64
	deletedFiles                []string
}

func (r *Reporter) increment(path string, bytes int64, match bool) {
	if match {
		r.removed += bytes
		r.deletedFiles = append(r.deletedFiles, path)
	} else {
		r.remaining += bytes
	}
	r.initial += bytes
}

func (r *Reporter) report() {
	fmt.Println("Success!\n")
	fmt.Printf("Before: %d kb\n", (r.initial)/1024)
	fmt.Printf("Removed: %d kb\n", r.removed/1024)
	fmt.Printf("Remaining: %d kb\n", r.remaining/1024)
	fmt.Printf("Deleted: %d files\n\n", len(r.deletedFiles))
}
