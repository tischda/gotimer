// +build windows

package main

import (
	"flag"
	"io/ioutil"
	"log"
)

var start string
var read string
var stop string
var clear bool
var verbose bool
var list bool

func init() {
	flag.StringVar(&start, "start", "REQUIRED", "start timer")
	flag.StringVar(&read, "read", "REQUIRED", "read timer (elapsed time)")
	flag.StringVar(&stop, "stop", "REQUIRED", "stop timer and print elapsed time")
	flag.BoolVar(&clear, "clear", false, "clear all timers")
	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.BoolVar(&list, "list", false, "list timers")
}

// TODO: -C"..."  execute the command-line and display the execution time
// TODO: -list option, to list all timers

func main() {
	registry = realRegistry{}

	// configure logging
	log.SetFlags(0)

	// parse command line arguments
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
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
}
