package registry

type MockRegistry struct {
	Timers map[string]uint64
}

func NewMockRegistry() MockRegistry {
	mock := MockRegistry{}
	mock.Timers = make(map[string]uint64)
	return mock
}

func (r MockRegistry) SetQword(path RegPath, valueName string, value uint64) error {
	r.Timers[valueName] = value
	return nil
}

func (r MockRegistry) GetQword(path RegPath, valueName string) (uint64, error) {
	return r.Timers[valueName], nil
}

func (r MockRegistry) DeleteValue(path RegPath, valueName string) error {
	delete(r.Timers, valueName)
	return nil
}

func (r MockRegistry) CreateKey(path RegPath) error {
	return nil
}

func (r MockRegistry) DeleteKey(path RegPath) error {
	return nil
}

func (r MockRegistry) EnumValues(path RegPath) ([]string, error) {
	keys := make([]string, 0, len(r.Timers))
	for k := range r.Timers {
		keys = append(keys, k)
	}
	return keys, nil
}
