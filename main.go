package main

import (
	"fmt"
	"log"
	"flag"
)

var start string
var stop string
var elapsed string
var clear bool

const (
	path = `SOFTWARE\Tischer`
	subkey = `timers`
)

func init() {
	flag.StringVar(&start, "start", "", "start timer")
	flag.StringVar(&stop, "stop", "", "stop timer")
	flag.StringVar(&elapsed, "elapsed", "", "print elapsed time for timer (do not stop)")
	flag.BoolVar(&clear, "clear", false, "clear all timers")
}

func main() {
	// parse command line parameters
	flag.Parse()

	// configure logging
	log.SetFlags(0)

	if clear {
		clearTimers()
	}

	// TODO: check flags, key is mandatory for start and stop
	if start != "" {
		setNanos(start)
	}

	if stop != "" {
		// get nanos and print result
		// clear timer

		// fmt.Println(getNanos("key1"))
		// fmt.Println(getNanos("key2"))
	}

	if elapsed != "" {
		// get nanos and print result
		// no NOT clear timer
	}
}

func getNanos(timer string) uint64 {
	nanos, err := registryGetQword(path + "\\" + subkey, timer)
	if (err != nil) {
		log.Fatalf("The timer %q has not been started, try `timer -start <key>`", timer)
	}
	return nanos
}

func setNanos(timer string) {
	createTimerGroup()
	log.Println("Starting timer", timer)
	// TODO: set value
}

func clearTimers() {
	deleteTimerGroup()
	fmt.Println("Timers deleted.")
}

func createTimerGroup() {
	err := registryCreateKey(path)
	if err != nil {
		log.Fatal(err)
	}
	err = registryCreateKey(path + "\\" + subkey)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteTimerGroup() {
	err := registryDeleteKey(path, subkey)
	if (err != nil) {
		log.Fatal(err)
	}
}
