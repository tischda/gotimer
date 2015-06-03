package main

type mockRegistry struct {
	timers map[string]uint64
}

func (r mockRegistry) SetQword(path string, valueName string, value uint64) error {
	r.timers[valueName] = value
	return nil
}

func (r mockRegistry) GetQword(path string, valueName string) (uint64, error) {
	return r.timers[valueName], nil
}

func (r mockRegistry) DeleteValue(path string, valueName string) error {
	delete(r.timers, valueName)
	return nil
}

func (r mockRegistry) CreateKey(path string) error {
	return nil
}

func (r mockRegistry) DeleteKey(path string, key string) error {
	return nil
}

func (r mockRegistry) EnumValues(path string) []string {
	keys := make([]string, 0, len(r.timers))
	for k := range r.timers {
		keys = append(keys, k)
	}
	return keys
}
