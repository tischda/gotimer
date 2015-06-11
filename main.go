// +build windows

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const version string = "1.1.1"

func main() {
	start := flag.String("start", "REQUIRED", "start timer")
	read := flag.String("read", "REQUIRED", "read timer (elapsed time)")
	stop := flag.String("stop", "REQUIRED", "stop timer and print elapsed time")
	command := flag.String("C", "REQUIRED", "print elapsed time for command")
	clear := flag.Bool("clear", false, "clear all timers / uninstall")
	verbose := flag.Bool("verbose", false, "verbose output")
	list := flag.Bool("list", false, "list timers")
	showVersion := flag.Bool("version", false, "print version and exit")

	// configure logging
	log.SetFlags(0)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] name\n  name: name of the timer\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if *showVersion {
		fmt.Println("timer version", version)
		return
	}

	t := timer{
		registry: realRegistry{},
	}

	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}

	if *clear {
		t.clearAll()
	}

	if *start != "REQUIRED" {
		t.start(*start)
	}

	if *read != "REQUIRED" {
		t.read(*read)
	}

	if *stop != "REQUIRED" {
		t.read(*stop)
		t.clear(*stop)
	}

	if *list {
		t.list()
	}

	// http://bit.ly/1dMD2YN
	if *command != "REQUIRED" {
		t.process("cmd", "/c", *command)
	}
}
