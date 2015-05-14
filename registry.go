package main

import (
	"log"
	"syscall"
	"unsafe"
)

func registrySetQword(path string, key string, value uint64) error {
	var handle syscall.Handle
	err := syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(path), 0, syscall.KEY_SET_VALUE, &handle)
	if err != nil {
		return err
	}
	defer syscall.RegCloseKey(handle)
	return regSetValueEx(handle, syscall.StringToUTF16Ptr(key), 0, syscall.REG_QWORD, (*byte)(unsafe.Pointer(&value)), 8)
}

func registryGetQword(path string, key string) (uint64, error) {
	var handle syscall.Handle
	err := syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(path), 0, syscall.KEY_READ, &handle)
	if err != nil {
		return 0, err
	}
	defer syscall.RegCloseKey(handle)

	var val uint64
	n := uint32(8)
	var typ uint32

	err = syscall.RegQueryValueEx(
		handle, syscall.StringToUTF16Ptr(key),
		nil,
		&typ,
		(*byte)(unsafe.Pointer(&val)),
		&n)

	if err != nil {
		return 0, err
	}

	if typ != syscall.REG_QWORD {
		log.Fatalln("Expected key of type REG_QWORD, but was", valueTypeName[typ])
	}
	return val, nil
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
