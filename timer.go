package main

import (
	"fmt"
	"log"
	"time"
)

const (
	path   = `SOFTWARE\Tischer`
	subkey = `timers`
)

// Starts the specified timer by creating a registry key containing
// the number of nanoseconds elapsed since January 1, 1970 UTC (int64).
func startTimer(timer string) {
	log.Println("Starting timer", timer)
	createTimersRegistryKey()
	registrySetQword(path+"\\"+subkey, timer, uint64(time.Now().UnixNano()))
}

// Prints the time elapsed since the timer record was created in the registry.
func readTimer(timer string) {
	fmt.Printf("Elapsed time (%s): %s\n", timer, getDuration(timer).String())
}

// Removes the timer record from the registry.
func clearTimer(timer string) {
	registryDeleteValue(path+"\\"+subkey, timer)
}

// Reads the timestamp recorded in the registry for this timer and
// calculates the duration from then to the current time.
func getDuration(timer string) time.Duration {
	nanos, err := registryGetQword(path+"\\"+subkey, timer)
	if err != nil {
		log.Fatalf("Timer record %q not found", timer)
	}
	t0 := time.Unix(0, int64(nanos))
	t1 := time.Now()
	return t1.Sub(t0)
}

// Removes all timers from the registry.
func clearAllTimers() {
	deleteTimersRegistryKey()
	log.Println("All timers deleted")
}

// Creates the timers subkey that will hold all timers. Note that
// if "path" does not exist, it will also be created.
func createTimersRegistryKey() {
	err := registryCreateKey(path + "\\" + subkey)
	if err != nil {
		log.Fatal(err)
	}
}

// Deletes the timers subkey that holds all timers. Doing this will
// effectively clear all timer records.
func deleteTimersRegistryKey() {
	registryDeleteKey(path, subkey)
}
