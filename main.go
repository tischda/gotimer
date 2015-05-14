package main

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"
)

const path = `SOFTWARE\Tischer\timers`

func main() {
	fmt.Println(getNanos("key1"))
}

// cf. ztypes_windows.go and https://msdn.microsoft.com/en-us/library/windows/desktop/ms724884(v=vs.85).aspx
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

func getNanos(key string) uint64 {
	var handle syscall.Handle
	err := syscall.RegOpenKeyEx(syscall.HKEY_CURRENT_USER, syscall.StringToUTF16Ptr(path), 0, syscall.KEY_READ, &handle)
	if err != nil {
		log.Fatalf("Cannot open path %q: %s", path, err)
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
        log.Fatalf("Cannot read key %q: %s", key, err)
	}

	if typ != syscall.REG_QWORD {
		log.Fatalln("Expected key of type REG_QWORD, but was", valueTypeName[typ])
	}
	return val
}
