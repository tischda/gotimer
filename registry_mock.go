package main

type mockRegistry struct {
	timers map[string]uint64
}

func (r mockRegistry) SetQword(key int, valueName string, value uint64) error {
	r.timers[valueName] = value
	return nil
}

func (r mockRegistry) GetQword(key int, valueName string) (uint64, error) {
	return r.timers[valueName], nil
}

func (r mockRegistry) DeleteValue(key int, valueName string) error {
	delete(r.timers, valueName)
	return nil
}

func (r mockRegistry) CreateKey(key int) error {
	return nil
}

func (r mockRegistry) DeleteKey(parent int, child int) error {
	return nil
}

func (r mockRegistry) EnumValues(key int) []string {
	keys := make([]string, 0, len(r.timers))
	for k := range r.timers {
		keys = append(keys, k)
	}
	return keys
}
