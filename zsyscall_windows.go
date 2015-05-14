package main

import (
	"syscall"
	"unsafe"
)

// cf. syscall/ztypes_windows.go and https://msdn.microsoft.com/en-us/library/windows/desktop/ms724884(v=vs.85).aspx
var valueTypeName = []string{
	syscall.REG_NONE:                       "REG_NONE",
	syscall.REG_SZ:                         "REG_SZ",
	syscall.REG_EXPAND_SZ:                  "REG_EXPAND_SZ",
	syscall.REG_BINARY:                     "REG_BINARY",
	syscall.REG_DWORD_LITTLE_ENDIAN:        "REG_DWORD_LITTLE_ENDIAN",
	syscall.REG_DWORD_BIG_ENDIAN:           "REG_DWORD_BIG_ENDIAN",
	syscall.REG_LINK:                       "REG_LINK",
	syscall.REG_MULTI_SZ:                   "REG_MULTI_SZ",
	syscall.REG_RESOURCE_LIST:              "REG_RESOURCE_LIST",
	syscall.REG_FULL_RESOURCE_DESCRIPTOR:   "REG_FULL_RESOURCE_DESCRIPTOR",
	syscall.REG_RESOURCE_REQUIREMENTS_LIST: "REG_RESOURCE_REQUIREMENTS_LIST",
	syscall.REG_QWORD_LITTLE_ENDIAN:        "REG_QWORD_LITTLE_ENDIAN",
}

// Adapted from https://github.com/golang/sys/blob/master/windows/registry/

var (
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")

	procRegCreateKeyExW = modadvapi32.NewProc("RegCreateKeyExW")
	procRegDeleteKeyW   = modadvapi32.NewProc("RegDeleteKeyW")
	procRegSetValueExW  = modadvapi32.NewProc("RegSetValueExW")
	procRegDeleteValueW = modadvapi32.NewProc("RegDeleteValueW")
)

func regCreateKeyEx(key syscall.Handle, subkey *uint16, reserved uint32, class *uint16, options uint32, desired uint32, sa *syscall.SecurityAttributes, result *syscall.Handle, disposition *uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall9(procRegCreateKeyExW.Addr(), 9, uintptr(key), uintptr(unsafe.Pointer(subkey)), uintptr(reserved), uintptr(unsafe.Pointer(class)), uintptr(options), uintptr(desired), uintptr(unsafe.Pointer(sa)), uintptr(unsafe.Pointer(result)), uintptr(unsafe.Pointer(disposition)))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func regDeleteKey(key syscall.Handle, subkey *uint16) (regerrno error) {
	r0, _, _ := syscall.Syscall(procRegDeleteKeyW.Addr(), 2, uintptr(key), uintptr(unsafe.Pointer(subkey)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func regSetValueEx(key syscall.Handle, valueName *uint16, reserved uint32, vtype uint32, buf *byte, bufsize uint32) (regerrno error) {
	r0, _, _ := syscall.Syscall6(procRegSetValueExW.Addr(), 6, uintptr(key), uintptr(unsafe.Pointer(valueName)), uintptr(reserved), uintptr(vtype), uintptr(unsafe.Pointer(buf)), uintptr(bufsize))
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}

func regDeleteValue(key syscall.Handle, name *uint16) (regerrno error) {
	r0, _, _ := syscall.Syscall(procRegDeleteValueW.Addr(), 2, uintptr(key), uintptr(unsafe.Pointer(name)), 0)
	if r0 != 0 {
		regerrno = syscall.Errno(r0)
	}
	return
}
