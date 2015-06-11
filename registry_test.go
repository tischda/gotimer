// +build windows

package main

import (
	"log"
	"testing"
)

func TestCreateOpenDeleteKey(t *testing.T) {

	var registry = realRegistry{}

	// create
	err := registry.CreateKey(PATH_TIMERS)
	if err != nil {
		log.Fatalln("Error in CreateKey", err)
	}

	// store value
	expected := uint64(1234)
	err = registry.SetQword(PATH_TIMERS, "t1", expected)
	if err != nil {
		log.Fatalln("Error in SetQword", err)
	}

	// list values
	timers1 := registry.EnumValues(PATH_TIMERS)
	if len(timers1) == 0 {
		t.Errorf("No timers found")
	}

	// read value
	actual, err1 := registry.GetQword(PATH_TIMERS, "t1")
	if err1 != nil {
		log.Fatalln("Error in GetQword", err1)
	}
	if actual != expected {
		t.Errorf("Expected: %q, was: %q", expected, actual)
	}

	// delete value
	err = registry.DeleteValue(PATH_TIMERS, "t1")
	if err != nil {
		t.Errorf("Error deleting value t1, %s", err)
	}
	timers2 := registry.EnumValues(PATH_TIMERS)
	if len(timers2) != len(timers1)-1 {
		t.Errorf("Timers should have been deleted")
	}

	// delete keys
	err = registry.DeleteKey(PATH_TIMERS)
	if err != nil {
		t.Errorf("Error deleting %s, %s", PATH_TIMERS.lpSubKey, err)
	}
	err = registry.DeleteKey(PATH_SOFTWARE)
	if err != nil {
		t.Errorf("Error deleting %s, %s", PATH_SOFTWARE.lpSubKey, err)
	}
}