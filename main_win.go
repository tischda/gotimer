// +build windows

package main

import "github.com/tischda/timer/registry"

func init() {
	shell = "cmd"
	shellCmdFlag = "/c"
	timer = theTimer{registry: registry.RealRegistry{}}
}
