package main

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var mock = mockRegistry{}

func init() {
	mock.timers = make(map[string]uint64)
	registry = mock
}

func TestStart(t *testing.T) {
	startTimer("t1")

	actual := mock.timers["t1"]
	if actual == 0 {
		t.Errorf("Expected: >0, was: %q", actual)
	}
}

func TestStop(t *testing.T) {
	startTimer("t2")
	time.Sleep(10 * time.Millisecond)
	actual := getDuration("t2")
	if actual < 9*time.Millisecond || actual > 14*time.Millisecond {
		t.Errorf("Expected: 10 msec, was: %q", actual)
	}
}

func TestClear(t *testing.T) {
	startTimer("t3")
	clearTimer("t3")
	_, exists := mock.timers["t3"]
	if exists {
		t.Errorf("Expected: false, was: %q", exists)
	}
}

func TestList(t *testing.T) {
	clearAllTimers()
	startTimer("t1")
	startTimer("t2")

	expected := "[t1 t2]\n"
	actual := captureOutput(listTimers)
	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}

func TestProcess(t *testing.T) {
	expected := "Time processing sleep: 1."
	actual := captureOutput(func() {
		process("sleep", "1")
	})[:25]
	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}

// captures Stdout and returns output of function f()
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
