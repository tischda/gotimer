package main

import (
    "fmt"
    "log"
    "syscall"
    "unsafe"
)

const path = `SOFTWARE\Microsoft\Windows NT\CurrentVersion`
const name = `ProductId`

func main() {
    var handle syscall.Handle

    err := syscall.RegOpenKeyEx(syscall.HKEY_LOCAL_MACHINE, syscall.StringToUTF16Ptr(path), 0, syscall.KEY_READ, &handle)
	if err != nil {
		log.Fatal(err)
	}
    defer syscall.RegCloseKey(handle)

    var buf[syscall.MAX_LONG_PATH]uint16
    n := uint32(len(buf))

    err = syscall.RegQueryValueEx(handle, syscall.StringToUTF16Ptr(name), nil, nil, (*byte)(unsafe.Pointer(&buf[0])), &n)
    if err != nil {
        return
    }

    fmt.Printf("%s=%q\n", name, syscall.UTF16ToString(buf[:]))
}
