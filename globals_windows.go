// +build windows

package main

import "github.com/tischda/timer/registry"

// The following variables are compiled for windows only

const shell = "cmd"
const shellCmdFlag = "/c"

// 'sleep' command available on appveyor but not in Windows
// TODO: write a sleeper that allows fractional numbers, eg. 0.1
// cf. http://golang.org/pkg/time/#Sleep
const execTestCmd = "sleep 1"
const execTestRxp = `Total time: 1.\d*s`

var timer Chronometer = &regTimer{registry: registry.RealRegistry{}}
