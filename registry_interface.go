package main

type Registry interface {
	SetQword(path string, valueName string, value uint64) error
	GetQword(path string, valueName string) (uint64, error)
	DeleteValue(path string, valueName string) error
	CreateKey(path string) error
	DeleteKey(path string, key string) error
	EnumValues(path string) []string
}
