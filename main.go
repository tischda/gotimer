package main

import (
	"flag"
	"log"
)

var start string
var read string
var stop string
var clear bool

func init() {
	flag.StringVar(&start, "start", "REQUIRED", "start timer")
	flag.StringVar(&read, "read", "REQUIRED", "read timer (elapsed time)")
	flag.StringVar(&stop, "stop", "REQUIRED", "stop timer and print elapsed time")
	flag.BoolVar(&clear, "clear", false, "clear all timers")
}

// TODO: -C"..."  execute the command-line and display the execution time
// TODO: -list option, to list all timers

func main() {
	// configure logging
	log.SetFlags(0)

	// parse command line arguments
	flag.Parse()

	// command line arguments are mandatory
	if flag.NFlag() == 0 {
		flag.Usage()
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
}
