package registry

import "testing"

var PATH_SOFTWARE = RegPath{HKEY_CURRENT_USER, `SOFTWARE\Tischer`}
var PATH_TIMERS = RegPath{HKEY_CURRENT_USER, `SOFTWARE\Tischer\timers`}

func TestCreateOpenDeleteKey(t *testing.T) {

	// create
	err := registry.CreateKey(PATH_TIMERS)
	if err != nil {
		t.Errorf("Error in CreateKey", err)
	}

	// store value
	expected := uint64(1234)
	err = registry.SetQword(PATH_TIMERS, "t1", expected)
	if err != nil {
		t.Errorf("Error in SetQword", err)
	}

	// list values
	timers1 := registry.EnumValues(PATH_TIMERS)
	if len(timers1) == 0 {
		t.Errorf("No timers found")
	}

	// read value
	actual, err1 := registry.GetQword(PATH_TIMERS, "t1")
	if err1 != nil {
		t.Errorf("Error in GetQword", err1)
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
		t.Errorf("Error deleting %s, %s", PATH_TIMERS.LpSubKey, err)
	}
	err = registry.DeleteKey(PATH_SOFTWARE)
	if err != nil {
		t.Errorf("Error deleting %s, %s", PATH_SOFTWARE.LpSubKey, err)
	}
}
