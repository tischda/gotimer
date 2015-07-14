// +build !windows

package main

import "github.com/tischda/timer/registry"

// The following variables are defined for non-windows machines
// allowing for basic testing.

const shell = "bash"
const shellCmdFlag = "-c"

// 'sleep' command available on appveyor but not in Windows
const execTestCmd = "sleep 0.1"
const execTestRxp = `Total time: 1\d\d.\d*ms`

var timer Chronometer = &regTimer{registry: registry.NewMockRegistry()}
