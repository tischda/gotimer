package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

// http://technosophos.com/2014/06/11/compile-time-string-in-go.html
// go build -ldflags "-x main.version=$(git describe --tags)"
var version string

// command line flags
var showVersion bool
var quietProcess bool

// command structure
type cliCmd struct {
	CmdName string
	CmdDesc string
	CmdFunc func(Chronometer, string)
}

const USAGE_TEMPLATE = `{{range .}}  {{.CmdName}}: {{.CmdDesc}}
{{end}}`

// list of permitted commands with description and func pointer
var cmdList = []cliCmd{
	{"start", "start named timer", Chronometer.start},
	{"read", "read timer (elapsed time)", Chronometer.read},
	{"stop", "read and then clear timer", Chronometer.stop},
	{"list", "list timers", Chronometer.list},
	{"clear", "clear named timer, remove from registry", Chronometer.clear},
	{"exec", "execute task and print elapsed time", Chronometer.exec},
}

// Note: defining flags in init() to avoid "panic: flag redefined"
// when calling main() twice in main_test.
func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.BoolVar(&quietProcess, "quiet", false, "hide process output")
}

func main() {
	log.SetFlags(0)
	flag.Usage = customUsage
	flag.Parse()

	if flag.Arg(0) == "version" || showVersion {
		fmt.Printf("gotimer %s - Measures time between two events\n", version)
	} else {
		processArgs(flag.Arg(0), flag.Arg(1))
	}
}

// Note: if flag not specified in first position, it will be ignored.
// Example: ./timer start azerty -C ls --> flag '-C' will be ignored
func processArgs(cmd string, name string) {
	checkParameters(cmd, name)
	executeTimerFunc(cmd, name)
}

// Returns a custom flag.Usage() function.
func customUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION] exec task\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "       %s [OPTION] COMMAND timer-name\n", os.Args[0])
	fmt.Fprint(os.Stderr, "\n COMMANDS:\n")
	tpl, _ := template.New("usage").Parse(USAGE_TEMPLATE)
	tpl.Execute(os.Stderr, cmdList)
	fmt.Fprint(os.Stderr, "\n OPTIONS:\n")
	flag.PrintDefaults()
	os.Exit(1)
}

// Verifies command line parameters.
func checkParameters(cmd string, name string) {
	if flag.NArg() == 0 || flag.NArg() > 2 {
		flag.Usage()
	}
	if (cmd == "start" || cmd == "stop") && name == "" {
		fmt.Fprint(os.Stderr, "Please specify the name of the timer.\n")
		os.Exit(1)
	}
	if cmd == "exec" && name == "" {
		fmt.Fprint(os.Stderr, "Please specify a task to execute.\n")
		os.Exit(1)
	}
}

// Executes the requested action defined in command list.
func executeTimerFunc(cmd string, name string) {
	i := indexOf(cmd)
	if i != -1 {
		cmdList[i].CmdFunc(timer, name)
	} else {
		flag.Usage()
	}
}

// Checks specified command against supported commands
// and return index if found, -1 otherwise.
func indexOf(cmdName string) int {
	for i, item := range cmdList {
		if item.CmdName == cmdName {
			return i
		}
	}
	return -1
}
