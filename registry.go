package main

import (
	"log"
	"syscall"
	"unsafe"
)

// Writes a REG_QWORD (uint64) to the Windows registry.
func registrySetQword(path string, valueName string, value uint64) error {
	handle := registryOpenKey(path, syscall.KEY_SET_VALUE)
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
func registryGetQword(path string, valueName string) (uint64, error) {
	handle := registryOpenKey(path, syscall.KEY_READ)
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
func registryDeleteValue(path string, valueName string) error {
	handle := registryOpenKey(path, syscall.KEY_WRITE)
	defer syscall.RegCloseKey(handle)
	return regDeleteValue(handle, syscall.StringToUTF16Ptr(valueName))
}

// Creates a key in the Windows registry.
func registryCreateKey(path string) error {

	// handle is required by function call, but not used
	var handle syscall.Handle

	// 1 - newly created
	// 2 - already existing
	var d uint32

	return regCreateKeyEx(
		syscall.HKEY_CURRENT_USER,
		syscall.StringToUTF16Ptr(path),
		0,
		nil,
		0,
		syscall.KEY_CREATE_SUB_KEY,
		nil,
		&handle,
		&d)
}

// Deletes a key from the Windows registry.
func registryDeleteKey(path string, key string) error {
	handle := registryOpenKey(path, syscall.KEY_WRITE)
	defer syscall.RegCloseKey(handle)
	return regDeleteKey(handle, syscall.StringToUTF16Ptr(key))
}

// Enumerates the values for the specified registry key. The function
// returns an array of valueNames.
func registryEnumValues(path string) []string {
	var values []string
	name, err := registryGetNextEnumValue(path, uint32(0))
	for i := 1; err == nil; i++ {
		values = append(values, name)
		name, err = registryGetNextEnumValue(path, uint32(i))
	}
	return values
}

// Enumerates the values for the specified registry key. The function
// returns one indexed value name for the key each time it is called.
func registryGetNextEnumValue(path string, index uint32) (string, error) {
	handle := registryOpenKey(path, syscall.KEY_READ)
	defer syscall.RegCloseKey(handle)

	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms724872(v=vs.85).aspx
	var nameLen uint32 = 16383
	name := make([]uint16, nameLen)

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

// Opens a Windows HKCU registry key and returns a handle. You must close
// the handle with `defer syscall.RegCloseKey(handle)` in the calling code.
func registryOpenKey(path string, desiredAccess uint32) syscall.Handle {
	var handle syscall.Handle
	err := syscall.RegOpenKeyEx(
		syscall.HKEY_CURRENT_USER,
		syscall.StringToUTF16Ptr(path),
		0,
		desiredAccess,
		&handle)
	if err != nil {
		log.Println("Cannot open path %q in the registry\n", path)
	}
	return handle
}
