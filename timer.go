package main

import (
	"fmt"
	"log"
	"time"
)

const (
	REG_PATH   string = `SOFTWARE\Tischer`
	REG_SUBKEY string = `timers`
)

// Starts the specified timer by creating a registry key containing
// the number of nanoseconds elapsed since January 1, 1970 UTC (int64).
func startTimer(timer string) {
	log.Println("Starting timer", timer)
	createTimersRegistryKey()
	// conversion int64 -> uint64 ok (nanos > 0)
	registrySetQword(REG_PATH+"\\"+REG_SUBKEY, timer, uint64(time.Now().UnixNano()))
}

// Prints the time elapsed since the timer record was created in the registry.
func readTimer(timer string) {
	fmt.Printf("Elapsed time (%s): %s\n", timer, getDuration(timer).String())
}

// Removes the timer record from the registry.
func clearTimer(timer string) {
	registryDeleteValue(REG_PATH+"\\"+REG_SUBKEY, timer)
}

// Reads the timestamp recorded in the registry for this timer and
// calculates the duration from then to the current time.
func getDuration(timer string) time.Duration {
	nanos, err := registryGetQword(REG_PATH+"\\"+REG_SUBKEY, timer)
	if err != nil {
		log.Fatalf("Timer record %q not found", timer)
	}
	// conversion uint64 -> int64 ok, since original value was int64
	t0 := time.Unix(0, int64(nanos))
	t1 := time.Now()
	return t1.Sub(t0)
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
	err := registryCreateKey(REG_PATH + "\\" + REG_SUBKEY)
	if err != nil {
		log.Fatal(err)
	}
}

// Deletes the timers subkey that holds all timers. Doing this will
// effectively clear all timer records.
func deleteTimersRegistryKey() {
	registryDeleteKey(REG_PATH, REG_SUBKEY)
}

func listTimers() {
	timers := registryEnumValues(REG_PATH + "\\" + REG_SUBKEY)
	if len(timers) == 0 {
		fmt.Println("No timers")
	} else {
		fmt.Println(timers)
	}
}
