package main

import (
	"fmt"
	"github.com/tischda/timer/registry"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
	"time"
)

var PATH_SOFTWARE = registry.RegPath{registry.HKEY_CURRENT_USER, `SOFTWARE\Tischer`}
var PATH_TIMERS = registry.RegPath{registry.HKEY_CURRENT_USER, `SOFTWARE\Tischer\timers`}

// Timer records time stamps in a registry
type Timer struct {
	registry registry.Registry
}

type Chronometer interface {
	start(name string)
	stop(name string)
	read(name string)
	clear(name string)
	list(name string)
	exec(process string)
}

// Starts the specified timer by creating a registry key containing
// the number of nanoseconds elapsed since January 1, 1970 UTC (int64).
func (t *Timer) start(name string) {
	checkFatal(t.registry.CreateKey(PATH_TIMERS))
	// conversion int64 -> uint64 ok (nanos > 0)
	checkFatal(t.registry.SetQword(PATH_TIMERS, name, uint64(time.Now().UnixNano())))
}

// Prints the time elapsed and removes the timer entry
func (t *Timer) stop(name string) {
	t.read(name)
	t.clear(name)
}

// Prints the time elapsed since the timer record was created in the registry.
func (t *Timer) read(name string) {
	fmt.Printf("Elapsed time (%s): %s\n", name, t.getDuration(name).String())
}

// Removes the timer entry from the registry.
// If none specified, clears all timers.
func (t *Timer) clear(name string) {
	if name != "" {
		checkFatal(t.registry.DeleteValue(PATH_TIMERS, name))
	} else {
		checkFatal(t.registry.DeleteKey(PATH_TIMERS))
		checkFatal(t.registry.DeleteKey(PATH_SOFTWARE))
	}
}

// List all started timers
func (t *Timer) list(name string) {
	timers := t.registry.EnumValues(PATH_TIMERS)
	if len(timers) > 0 {
		sort.Strings(timers)
		fmt.Println(timers)
	} else {
		fmt.Println("No timers.")
	}
}

// Execute process and print elapsed time
func (t *Timer) exec(process string) {
	defer whenDone()("Total time: %v\n")

	// execute process (http://bit.ly/1dMD2YN)
	cmd := exec.Command(shell, shellCmdFlag, process)
	if !quietProcess {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	checkFatal(cmd.Run())
}

// Reads the timestamp recorded in the registry for this timer and
// calculates the duration from then to the current time.
func (t *Timer) getDuration(name string) time.Duration {
	t1 := time.Now()
	nanos, err := t.registry.GetQword(PATH_TIMERS, name)
	checkFatal(err)
	// conversion uint64 -> int64 ok, since original value was int64
	t0 := time.Unix(0, int64(nanos))
	return t1.Sub(t0)
}

// Callback function executed when process is done
func whenDone() func(format string, args ...interface{}) {
	start := time.Now()
	return func(format string, args ...interface{}) {
		fmt.Printf(format, append(args, time.Since(start))...)
	}
}

// Print error and exit if err != nil
func checkFatal(err error) {
	if err != nil {
		message := err.Error()
		if strings.Contains(message, "The system cannot find the file specified.") {
			log.Fatal("The system cannot find the timer specified.")
		} else {
			log.Fatalln(err)
		}
	}
}
