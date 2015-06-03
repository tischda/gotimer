package main

import (
	"fmt"
	"log"
	"os/exec"
	"sort"
	"time"
)

const (
	REG_PATH   string = `SOFTWARE\Tischer`
	REG_SUBKEY string = `timers`
)

var registry Registry

// Starts the specified timer by creating a registry key containing
// the number of nanoseconds elapsed since January 1, 1970 UTC (int64).
func startTimer(timer string) {
	log.Println("Starting timer", timer)
	createTimersRegistryKey()
	// conversion int64 -> uint64 ok (nanos > 0)
	registry.SetQword(REG_PATH+"\\"+REG_SUBKEY, timer, uint64(time.Now().UnixNano()))
}

// Prints the time elapsed since the timer record was created in the registry.
func readTimer(timer string) {
	fmt.Printf("Elapsed time (%s): %s\n", timer, getDuration(timer).String())
}

// Removes the timer record from the registry.
func clearTimer(timer string) {
	registry.DeleteValue(REG_PATH+"\\"+REG_SUBKEY, timer)
}

// Reads the timestamp recorded in the registry for this timer and
// calculates the duration from then to the current time.
func getDuration(timer string) time.Duration {
	nanos, err := registry.GetQword(REG_PATH+"\\"+REG_SUBKEY, timer)
	if err != nil {
		log.Fatalf("Timer record %q not found", timer)
	}
	// conversion uint64 -> int64 ok, since original value was int64
	start := time.Unix(0, int64(nanos))
	return time.Since(start)
}

// Removes all timers from the registry.
func clearAllTimers() {
	deleteTimersRegistryKey()
	createTimersRegistryKey()
	log.Println("All timers deleted")
}

// Creates the timers subkey that will hold all timers. Note that
// if "path" does not exist, it will also be created.
func createTimersRegistryKey() {
	err := registry.CreateKey(REG_PATH + "\\" + REG_SUBKEY)
	if err != nil {
		log.Fatal(err)
	}
}

// Deletes the timers subkey that holds all timers. Doing this will
// effectively clear all timer records.
func deleteTimersRegistryKey() {
	registry.DeleteKey(REG_PATH, REG_SUBKEY)
}

func listTimers() {
	timers := registry.EnumValues(REG_PATH + "\\" + REG_SUBKEY)
	if len(timers) == 0 {
		fmt.Println("No timers")
	} else {
		sort.Strings(timers)
		fmt.Println(timers)
	}
}

func process(name string, args ...string) {
	defer whenDone()("Time processing %v: %v", name)
	exec.Command(name, args...).Run()
}

func whenDone() func(format string, args ...interface{}) {
	start := time.Now()
	return func(format string, args ...interface{}) {
		fmt.Printf(format, append(args, time.Since(start))...)
	}
}
