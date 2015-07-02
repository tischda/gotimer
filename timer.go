package main

import (
	"fmt"
	"github.com/tischda/timer/registry"
	"log"
	"os"
	"os/exec"
	"sort"
	"time"
)

type timer struct {
	registry registry.Registry
}

var PATH_SOFTWARE = registry.RegPath{registry.HKEY_CURRENT_USER, `SOFTWARE\Tischer`}
var PATH_TIMERS = registry.RegPath{registry.HKEY_CURRENT_USER, `SOFTWARE\Tischer\timers`}

// Starts the specified timer by creating a registry key containing
// the number of nanoseconds elapsed since January 1, 1970 UTC (int64).
func (t *timer) start(name string) {
	log.Println("Starting timer", name)
	checkFatal(t.registry.CreateKey(PATH_TIMERS))
	// conversion int64 -> uint64 ok (nanos > 0)
	checkFatal(t.registry.SetQword(PATH_TIMERS, name, uint64(time.Now().UnixNano())))
}

// Prints the time elapsed since the timer record was created in the registry.
func (t *timer) read(name string) {
	fmt.Printf("Elapsed time (%s): %s\n", name, t.getDuration(name).String())
}

// Removes the timer entries from the registry.
func (t *timer) clear(name string) {
	checkFatal(t.registry.DeleteValue(PATH_TIMERS, name))
}

// Reads the timestamp recorded in the registry for this timer and
// calculates the duration from then to the current time.
func (t *timer) getDuration(name string) time.Duration {
	nanos, err := t.registry.GetQword(PATH_TIMERS, name)
	checkFatal(err)
	// conversion uint64 -> int64 ok, since original value was int64
	start := time.Unix(0, int64(nanos))
	return time.Since(start)
}

// Removes all timers keys from the registry.
func (t *timer) clearAll() {
	checkFatal(t.registry.DeleteKey(PATH_TIMERS))
	checkFatal(t.registry.DeleteKey(PATH_SOFTWARE))
}

func (t *timer) list() {
	timers := t.registry.EnumValues(PATH_TIMERS)
	if len(timers) > 0 {
		sort.Strings(timers)
		fmt.Println(timers)
	}
}

func (t *timer) process(name string, args ...string) {
	defer whenDone()("Total time: %v\n")
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func whenDone() func(format string, args ...interface{}) {
	start := time.Now()
	return func(format string, args ...interface{}) {
		fmt.Printf(format, append(args, time.Since(start))...)
	}
}

func checkFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
