// +build windows

package main

import (
	"os"
	"testing"
)

func TestMainTimer(t *testing.T) {
	args := []string{"-start", "tmain", "-stop", "tmain"}
	os.Args = append(os.Args, args...)

	expected := "Elapsed time (tmain): 0\n"

	// this can be done only once or test framework will panic
	actual := captureOutput(main)

	if expected != actual {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}
