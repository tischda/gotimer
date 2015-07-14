package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

// Inspired by https://talks.golang.org/2014/testing.slide#23
func TestUsage(t *testing.T) {

	if os.Getenv("BE_CRASHER") == "1" {
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestUsage")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")

	// capture output of process execution
	r, w, _ := os.Pipe()
	cmd.Stderr = w
	err := cmd.Run()
	w.Close()

	// check return code
	if e, ok := err.(*exec.ExitError); ok && e.Success() {
		t.Fatalf("Exptected exit status 1, but was: %v, ", err)
	}

	// now check that Usage message is displayed
	captured, _ := ioutil.ReadAll(r)
	actual := string(captured)
	expected := "Usage:"

	if !strings.Contains(actual, expected) {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}

// Tests that specified actions call corresponding functions.
func TestParams(t *testing.T) {
	mockTimer := mockTimer{}
	timer = &mockTimer
	init_args := os.Args

	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"start", "name"}, "start"},
		{[]string{"read", "name"}, "read"},
		{[]string{"stop", "name"}, "stop"},
		{[]string{"clear"}, "clear"},
		{[]string{"list"}, "list"},
		{[]string{"exec", "process"}, "process"},
	}

	for _, c := range cases {
		os.Args = append(init_args, c.in...)
		main()
		switch c.want {
		case "start":
			if !mockTimer.startCalled {
				t.Error("Timer expected to start but did not")
			}

		case "read":
			if !mockTimer.readCalled {
				t.Error("Timer expected to read but did not")
			}

		case "stop":
			if !mockTimer.stopCalled {
				t.Error("Timer expected to stop but did not")
			}

		case "clear":
			if !mockTimer.clearCalled {
				t.Error("Timer expected to clear but did not")
			}
		case "list":
			if !mockTimer.listCalled {
				t.Error("Timer expected to list but did not")
			}
		case "exec":
			if !mockTimer.execCalled {
				t.Error("Timer expected to exec but did not")
			}
		}
	}
}

func TestIndexOf(t *testing.T) {
	expected := 4
	actual := indexOf("clear")
	assertEquals(t, expected, actual)

	expected = -1
	actual = indexOf("toto")
	assertEquals(t, expected, actual)
}

func assertEquals(t *testing.T, expected int, actual int) {
	if actual != expected {
		t.Errorf("Expected: %s, but was: %s", expected, actual)
	}
}
