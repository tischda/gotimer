package main

type mockRegistry struct {
	timers map[string]uint64
}

var mock = mockRegistry{}

func init() {
	mock.timers = make(map[string]uint64)
}

func (r mockRegistry) SetQword(path regPath, valueName string, value uint64) error {
	r.timers[valueName] = value
	return nil
}

func (r mockRegistry) GetQword(path regPath, valueName string) (uint64, error) {
	return r.timers[valueName], nil
}

func (r mockRegistry) DeleteValue(path regPath, valueName string) error {
	delete(r.timers, valueName)
	return nil
}

func (r mockRegistry) CreateKey(path regPath) error {
	return nil
}

func (r mockRegistry) DeleteKey(path regPath) error {
	return nil
}

func (r mockRegistry) EnumValues(path regPath) []string {
	keys := make([]string, 0, len(r.timers))
	for k := range r.timers {
		keys = append(keys, k)
	}
	return keys
}
