package main

import (
	"testing"
)

func TestSplit(t *testing.T) {
	expected_path := regPath{HKEY_CURRENT_USER, `SOFTWARE\Tischer`}
	expected_key := `timers`

	actual_path, actual_key := splitPathSubkey(PATH_TIMERS)
	if actual_path != expected_path {
		t.Errorf("Expected: %q, was: %q", expected_path, actual_path)
	}
	if actual_key != expected_key {
		t.Errorf("Expected: %q, was: %q", expected_key, actual_key)
	}
}
