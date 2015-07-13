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

var PATH_SOFTWARE = registry.RegPath{registry.HKEY_CURRENT_USER, `SOFTWARE\Tischer`}
var PATH_TIMERS = registry.RegPath{registry.HKEY_CURRENT_USER, `SOFTWARE\Tischer\timers`}

// TODO: naming of type and interface, what is best practice ?
type theTimer struct {
	registry registry.Registry
}

type Timer interface {
	start(name string)
	stop(name string)
	read(name string)
	clear(name string)
	list(name string)
	exec(process string)
}

// Starts the specified timer by creating a registry key containing
// the number of nanoseconds elapsed since January 1, 1970 UTC (int64).
func (t *theTimer) start(name string) {
	checkFatal(t.registry.CreateKey(PATH_TIMERS))
	// conversion int64 -> uint64 ok (nanos > 0)
	checkFatal(t.registry.SetQword(PATH_TIMERS, name, uint64(time.Now().UnixNano())))
}

// Prints the time elapsed and removes the timer entry
func (t *theTimer) stop(name string) {
	t.read(name)
	t.clear(name)
}

// Prints the time elapsed since the timer record was created in the registry.
func (t *theTimer) read(name string) {
	fmt.Printf("Elapsed time (%s): %s\n", name, t.getDuration(name).String())
}

// Removes the timer entries from the registry.
func (t *theTimer) clear(name string) {
	checkFatal(t.registry.DeleteValue(PATH_TIMERS, name))
}

func (t *theTimer) list(name string) {
	timers := t.registry.EnumValues(PATH_TIMERS)
	if len(timers) > 0 {
		sort.Strings(timers)
		fmt.Println(timers)
	}
}

func (t *theTimer) exec(process string) {
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

// TODO: small delay added by registry lookup. Should stop timer immediately
// and then lookup.
func (t *theTimer) getDuration(name string) time.Duration {
	nanos, err := t.registry.GetQword(PATH_TIMERS, name)
	checkFatal(err)
	// conversion uint64 -> int64 ok, since original value was int64
	start := time.Unix(0, int64(nanos))
	return time.Since(start)
}

// Removes all timers keys from the registry.
func (t *theTimer) clearAll() {
	checkFatal(t.registry.DeleteKey(PATH_TIMERS))
	checkFatal(t.registry.DeleteKey(PATH_SOFTWARE))
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
		log.Fatalln(err)
	}
}
