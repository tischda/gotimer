// +build windows

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"fmt"
)

var start string
var read string
var stop string
var clear bool
var verbose bool
var list bool
var version bool
var command string

func init() {
	flag.StringVar(&start, "start", "REQUIRED", "start timer")
	flag.StringVar(&read, "read", "REQUIRED", "read timer (elapsed time)")
	flag.StringVar(&stop, "stop", "REQUIRED", "stop timer and print elapsed time")
	flag.StringVar(&command, "C", "REQUIRED", "print elapsed time for command")
	flag.BoolVar(&clear, "clear", false, "clear all timers")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.BoolVar(&version, "version", false, "print version")
	flag.BoolVar(&list, "list", false, "list timers")
}

func main() {
	registry = realRegistry{}

	// configure logging
	log.SetFlags(0)

	// parse command line arguments
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
	}

	if version {
		fmt.Println("Timer v1.1.0")
		return
	}

	if !verbose {
		log.SetOutput(ioutil.Discard)
	}

	if clear {
		clearAllTimers()
	}

	if start != "REQUIRED" {
		startTimer(start)
	}

	if read != "REQUIRED" {
		readTimer(read)
	}

	if stop != "REQUIRED" {
		readTimer(stop)
		clearTimer(stop)
	}

	if list {
		listTimers()
	}

	// http://bit.ly/1dMD2YN
	if command != "REQUIRED" {
		process("cmd", "/c", command)
	}
}
