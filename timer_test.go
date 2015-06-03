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

	// redirect output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	listTimers()

	// reset output again
	w.Close()
	os.Stdout = old

	captured, _ := ioutil.ReadAll(r)

	expected := "[t1 t2]\n"
	actual := string(captured)
	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}
