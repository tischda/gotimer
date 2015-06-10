package main

// Registry key indexes, do not reorder
const (
	SOFTWARE = iota
	TIMERS
	TIMERS_CHILD
)

var registry Registry

type Registry interface {
	SetQword(key int, valueName string, value uint64) error
	GetQword(key int, valueName string) (uint64, error)
	DeleteValue(key int, valueName string) error
	CreateKey(key int) error
	DeleteKey(parent int, child int) error
	EnumValues(key int) []string
}
