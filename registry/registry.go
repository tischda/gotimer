// Package registry provides primitives to access the Windows Registry
package registry

// A registry path is composed of an hKey index and the string representation
// of the path withing that hKey. We use hKey indexes to avoid dependency on
// non-portable syscall values.
type RegPath struct {
	HKeyIdx  uint8
	LpSubKey string
}

// Registry hKey index values, do not reorder
const (
	HKEY_CLASSES_ROOT = iota
	HKEY_CURRENT_USER
	HKEY_LOCAL_MACHINE
	HKEY_USERS
	HKEY_PERFORMANCE_DATA
	HKEY_CURRENT_CONFIG
	HKEY_DYN_DATA
)

type Registry interface {
	SetQword(path RegPath, valueName string, value uint64) error
	GetQword(path RegPath, valueName string) (uint64, error)
	DeleteValue(path RegPath, valueName string) error
	CreateKey(path RegPath) error
	DeleteKey(path RegPath) error
	EnumValues(ath RegPath) []string
}
