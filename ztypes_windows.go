// +build windows

package main

import (
	"syscall"
)

// Type description for verbose error messages. See:
// * syscall/ztypes_windows.go
// * https://msdn.microsoft.com/en-us/library/windows/desktop/ms724884(v=vs.85).aspx
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
