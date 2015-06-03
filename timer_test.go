package main

import (
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
	if actual < 9 * time.Millisecond || actual > 14 * time.Millisecond {
		t.Errorf("Expected: 10 msec, was: %q", actual)
	}
}
