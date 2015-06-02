package main

import (
	"testing"
)

var mock = mockRegistry{}

func init() {
	mock.timers = make(map[string]uint64)
	registry = mock
}

func TestHello(t *testing.T) {
	startTimer("toto")

	actual := mock.timers["toto"]
	if actual == 0 {
		t.Errorf("Expected: >0, was: %q", actual)
	}
}
