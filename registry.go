// +build windows

package main

import (
	"log"
	"syscall"
	"unsafe"
)

type realRegistry struct{}

// do not reorder
var hKeyTable = []syscall.Handle{
	syscall.HKEY_CLASSES_ROOT,
	syscall.HKEY_CURRENT_USER,
	syscall.HKEY_LOCAL_MACHINE,
	syscall.HKEY_USERS,
	syscall.HKEY_PERFORMANCE_DATA,
	syscall.HKEY_CURRENT_CONFIG,
	syscall.HKEY_DYN_DATA,
}

// Writes a REG_QWORD (uint64) to the Windows registry.
func (realRegistry) SetQword(path regPath, valueName string, value uint64) error {
	handle := openKey(path, syscall.KEY_SET_VALUE)
	defer syscall.RegCloseKey(handle)

	return regSetValueEx(
		handle,
		syscall.StringToUTF16Ptr(valueName),
		0,
		syscall.REG_QWORD,
		(*byte)(unsafe.Pointer(&value)),
		8)
}

// Reads a REG_QWORD (uint64) from the Windows registry.
func (realRegistry) GetQword(path regPath, valueName string) (uint64, error) {
	handle := openKey(path, syscall.KEY_READ)
	defer syscall.RegCloseKey(handle)

	var value uint64
	n := uint32(8)
	var vtype uint32

	err := syscall.RegQueryValueEx(
		handle,
		syscall.StringToUTF16Ptr(valueName),
		nil,
		&vtype,
		(*byte)(unsafe.Pointer(&value)),
		&n)

	if err != nil {
		return 0, err
	}

	if vtype != syscall.REG_QWORD {
		log.Fatalln("Expected key of type REG_QWORD, but was", valueTypeName[vtype])
	}
	return value, nil
}

// Deletes a key value from the Windows registry.
func (realRegistry) DeleteValue(path regPath, valueName string) error {
	handle := openKey(path, syscall.KEY_WRITE)
	defer syscall.RegCloseKey(handle)

	return regDeleteValue(handle, syscall.StringToUTF16Ptr(valueName))
}

// Creates a key in the Windows registry.
func (realRegistry) CreateKey(path regPath) error {

	// handle is required by function call, but not used
	var handle syscall.Handle

	// 1 - newly created
	// 2 - already existing
	var d uint32

	return regCreateKeyEx(
		hKeyTable[path.hKeyIdx],
		syscall.StringToUTF16Ptr(path.lpSubKey),
		0,
		nil,
		0,
		syscall.KEY_CREATE_SUB_KEY,
		nil,
		&handle,
		&d)
}

// Deletes a key from the Windows registry.  All sub-keys must be
// deleted before deleting the key, or you will get `access denied`.
func (realRegistry) DeleteKey(path regPath) error {
	parent, child := splitPathSubkey(path)

	handle := openKey(parent, syscall.KEY_WRITE)
	defer syscall.RegCloseKey(handle)

	return regDeleteKey(handle, syscall.StringToUTF16Ptr(child))
}

// Enumerates the values for the specified registry key index. The function
// returns an array of valueNames.
func (realRegistry) EnumValues(path regPath) []string {
	var values []string
	name, err := getNextEnumValue(path, uint32(0))
	for i := 1; err == nil; i++ {
		values = append(values, name)
		name, err = getNextEnumValue(path, uint32(i))
	}
	return values
}

// Enumerates the values for the specified registry key. The function
// returns one indexed value name for the key each time it is called.
func getNextEnumValue(path regPath, index uint32) (string, error) {
	handle := openKey(path, syscall.KEY_READ)
	defer syscall.RegCloseKey(handle)

	var nameLen uint32 = 16383
	name := make([]uint16, nameLen)

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724872(v=vs.85).aspx
	err := regEnumValue(
		handle,
		index,
		&name[0],
		&nameLen,
		nil,
		nil,
		nil,
		nil)

	return syscall.UTF16ToString(name), err
}

// Opens a Windows registry key and returns a handle. You must close
// the handle with `defer syscall.RegCloseKey(handle)` in the calling code.
func openKey(path regPath, desiredAccess uint32) syscall.Handle {
	var handle syscall.Handle

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724897(v=vs.85).aspx
	err := syscall.RegOpenKeyEx(
		hKeyTable[path.hKeyIdx],
		syscall.StringToUTF16Ptr(path.lpSubKey),
		0,
		desiredAccess,
		&handle)

	if err != nil {
		log.Fatalln("Cannot open registry path:", path)
	}
	return handle
}
