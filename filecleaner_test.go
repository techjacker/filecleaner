package main

import (
	"fmt"
	"os"
	"os/exec"
	// "path/filepath"
	"testing"
)

var tmp string = os.TempDir() + "filecleaner"
var fixtures string = relPath("fixtures")

/*--------------------------------------
teardowns
---------------------------------------*/
func Test_Clean(t *testing.T) {
	before()
	testClean(t)
	after()
}

/*--------------------------------------
test
---------------------------------------*/
func testClean(t *testing.T) {
	const expectedDeleted = 8

	Clean(tmp)

	if len(reporter.deletedFiles) != expectedDeleted {
		fmt.Printf("Clean(%s) should have deleted %v files", tmp, expectedDeleted)
		t.Errorf("Clean(%s) should have deleted %v files", tmp, expectedDeleted)
	}
}

/*--------------------------------------
fixtures
---------------------------------------*/
func before() {
	after()
	createFixtures(fixtures, tmp)
}

func after() {
	os.Remove(tmp)
}

func createFixtures(src string, dest string) {
	if err := cp(src, dest); err != nil {
		fmt.Println(err)
	}
}

func cp(src string, dest string) error {
	cpCmd := exec.Command("cp", "-rf", src, dest)
	return cpCmd.Run()
}
