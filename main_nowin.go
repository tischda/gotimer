// -build windows

package main

import "github.com/tischda/timer/registry"

func init() {
	shell = "bash"
	shellCmdFlag = "-c"
	timer = &theTimer{registry: registry.NewMockRegistry()}
}
