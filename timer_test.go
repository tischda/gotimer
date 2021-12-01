package main

import (
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/tischda/gotimer/registry"
)

var sut Timer
var mockRegistry = registry.NewMockRegistry()

func init() {
	sut = Timer{
		registry: mockRegistry,
	}
}

func TestStart(t *testing.T) {
	sut.start("t1")
	actual := mockRegistry.Timers["t1"]
	if actual == 0 {
		t.Errorf("Expected: >0, was: %q", actual)
	}
}

func TestStop(t *testing.T) {
	sut.start("t2")
	time.Sleep(10 * time.Millisecond)
	actual := sut.getDuration("t2")
	// This fails quite often in AppVeyor...
	if actual < 9*time.Millisecond || actual > 30*time.Millisecond {
		t.Errorf("Expected: ~10 msec, was: %q", actual)
	}
}

func TestClear(t *testing.T) {
	sut.start("t3")
	sut.clear("t3")
	_, exists := mockRegistry.Timers["t3"]
	if exists {
		t.Errorf("Expected: false, was: %q", strconv.FormatBool(exists))
	}
}

func TestList(t *testing.T) {
	sut.clear("")
	sut.start("t1")
	sut.start("t2")

	expected := "[t1 t2]\n"
	actual := captureOutput(func() { sut.list("") })
	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}

func TestExec(t *testing.T) {
	r, _ := regexp.Compile(execTestRxp)
	actual := captureOutput(func() { sut.exec(execTestCmd) })
	if !r.MatchString(actual) {
		t.Errorf("Expected: %q, was: %q", execTestRxp, actual)
	}
}

// Captures Stdout and returns output of function f()
func captureOutput(f func()) string {
	// redirect output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	// reset output again
	w.Close()
	os.Stdout = old

	captured, _ := ioutil.ReadAll(r)
	return string(captured)
}
