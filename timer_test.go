package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

var sut timer
var mock mockRegistry

func init() {
	mock = mockRegistry{}
	mock.timers = make(map[string]uint64)

	sut = timer{
		registry: mock,
	}
	log.SetOutput(ioutil.Discard)
}

func TestStart(t *testing.T) {
	sut.start("t1")
	actual := mock.timers["t1"]
	if actual == 0 {
		t.Errorf("Expected: >0, was: %q", actual)
	}
}

func TestStop(t *testing.T) {
	sut.start("t2")
	time.Sleep(10 * time.Millisecond)
	actual := sut.getDuration("t2")
	if actual < 9*time.Millisecond || actual > 14*time.Millisecond {
		t.Errorf("Expected: 10 msec, was: %q", actual)
	}
}

func TestClear(t *testing.T) {
	sut.start("t3")
	sut.clear("t3")
	_, exists := mock.timers["t3"]
	if exists {
		t.Errorf("Expected: false, was: %q", exists)
	}
}

func TestList(t *testing.T) {
	sut.clearAll()
	sut.start("t1")
	sut.start("t2")

	expected := "[t1 t2]\n"
	actual := captureOutput(sut.list)
	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}
}

func TestProcess(t *testing.T) {
	expected := "Total time: 1."
	actual := captureOutput(func() {
		sut.process("sleep", "1")
	})[:len(expected)]
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
