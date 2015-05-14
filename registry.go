package main

import (
	"log"
	"syscall"
	"unsafe"
)

func registrySetQword(path string, valueName string, value uint64) error {
	var handle syscall.Handle
	err := syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(path), 0, syscall.KEY_SET_VALUE, &handle)
	if err != nil {
		return err
	}
	defer syscall.RegCloseKey(handle)
	return regSetValueEx(handle, syscall.StringToUTF16Ptr(valueName), 0, syscall.REG_QWORD, (*byte)(unsafe.Pointer(&value)), 8)
}

func registryGetQword(path string, valueName string) (uint64, error) {
	var handle syscall.Handle
	err := syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(path), 0, syscall.KEY_READ, &handle)
	if err != nil {
		return 0, err
	}
	defer syscall.RegCloseKey(handle)

	var value uint64
	n := uint32(8)
	var vtype uint32

	err = syscall.RegQueryValueEx(
		handle, syscall.StringToUTF16Ptr(valueName),
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

func registryDeleteValue(path string, valueName string) error {
	var handle syscall.Handle
	err := syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(path), 0, syscall.KEY_WRITE, &handle)
	if err != nil {
		return err
	}
	defer syscall.RegCloseKey(handle)
	return regDeleteValue(handle, syscall.StringToUTF16Ptr(valueName))
}

func registryCreateKey(path string) error {
	var handle syscall.Handle

	// 1 - newly created
	// 2 - already existing
	var d uint32

	err := regCreateKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(path),
		0, nil, 0, syscall.KEY_CREATE_SUB_KEY, nil, &handle, &d)

	return err
}

func registryDeleteKey(path string, key string) error {
	var handle syscall.Handle
	err := syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(path), 0, syscall.KEY_WRITE, &handle)
	if err != nil {
		return err
	}
	defer syscall.RegCloseKey(handle)
	return regDeleteKey(handle, syscall.StringToUTF16Ptr(key))
}
