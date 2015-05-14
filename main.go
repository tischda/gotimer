package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var start string
var stop string
var elapsed string
var clear bool

const (
	path   = `SOFTWARE\Tischer`
	subkey = `timers`
)

func init() {
	flag.StringVar(&start, "start", "REQUIRED", "start timer")
	flag.StringVar(&stop, "stop", "REQUIRED", "stop timer and print elapsed time")
	flag.StringVar(&elapsed, "elapsed", "REQUIRED", "print elapsed time for timer")
	flag.BoolVar(&clear, "clear", false, "clear all timers")
}

// TODO: -C"..."  execute the command-line and display the execution time
// TODO: -list option, to list all timers

func main() {
	flag.Parse()
	log.SetFlags(0)

	if flag.NFlag() == 0 {
		flag.Usage()
	}

	if clear {
		clearAllTimers()
	}

	if start != "REQUIRED" {
		setNanos(start)
	}

	if stop != "REQUIRED" {
		printElapsed(stop)
		clearTimer(stop)
	}

	if elapsed != "REQUIRED" {
		printElapsed(elapsed)
	}
}

func printElapsed(timer string) {
	t0 := time.Unix(0, int64(getNanos(timer)))
	t1 := time.Now()
	duration := t1.Sub(t0)
	fmt.Printf("Elapsed time (%s): %s\n", timer, duration.String())
}

func getNanos(timer string) uint64 {
	nanos, err := registryGetQword(path+"\\"+subkey, timer)
	if err != nil {
		log.Fatalf("The timer %q has not been started", timer)
	}
	return nanos
}

func setNanos(timer string) {
	createTimersRegistryKey()
	log.Println("Starting timer", timer)
	registrySetQword(path+"\\"+subkey, timer, uint64(time.Now().UnixNano()))
}

func clearTimer(timer string) {
	registryDeleteValue(path+"\\"+subkey, timer)
}

func clearAllTimers() {
	deleteTimersRegistryKey()
	fmt.Println("All timers deleted")
}

// If "path" does not exist, it will be created
func createTimersRegistryKey() {
	err := registryCreateKey(path + "\\" + subkey)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteTimersRegistryKey() {
	err := registryDeleteKey(path, subkey)
	if err != nil {
		log.Fatal(err)
	}
}
