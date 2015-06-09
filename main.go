// +build windows

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const version string = "1.1.0"

func main() {
	start := flag.String("start", "REQUIRED", "start timer")
	read := flag.String("read", "REQUIRED", "read timer (elapsed time)")
	stop := flag.String("stop", "REQUIRED", "stop timer and print elapsed time")
	command := flag.String("C", "REQUIRED", "print elapsed time for command")
	clear := flag.Bool("clear", false, "clear all timers")
	verbose := flag.Bool("verbose", false, "verbose output")
	list := flag.Bool("list", false, "list timers")
	showVersion := flag.Bool("version", false, "print version")

	registry = realRegistry{}

	// configure logging
	log.SetFlags(0)

	// parse command line arguments
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] name\n  name: name of the timer\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("coucou")

	if *showVersion {
		fmt.Print("timer version", version)
		return
	}

	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}

	if *clear {
		clearAllTimers()
	}

	if *start != "REQUIRED" {
		startTimer(*start)
	}

	if *read != "REQUIRED" {
		readTimer(*read)
	}

	if *stop != "REQUIRED" {
		readTimer(*stop)
		clearTimer(*stop)
	}

	if *list {
		listTimers()
	}

	// http://bit.ly/1dMD2YN
	if *command != "REQUIRED" {
		process("cmd", "/c", *command)
	}
}
