package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

const version string = "1.2.0"

var shell, shellCmdFlag string

var timer Timer

var showVersion bool
var quietProcess bool

var stdout = os.Stdout
var stderr = os.Stderr

type cliCmd struct {
	CmdName string
	CmdDesc string
	CmdFunc func(Timer, string)
}

var cmdList = []cliCmd{
	cliCmd{"start", "start timer", Timer.start},
	cliCmd{"read", "read timer (elapsed time)", Timer.read},
	cliCmd{"stop", "read and then clear timer", Timer.stop},
	cliCmd{"list", "list timers", Timer.list},
	cliCmd{"clear", "clear all timers", Timer.clear},
	cliCmd{"exec", "execute process and print elapsed time", Timer.exec},
}

// define flags in init() otherwise you'll get "panic: flag redefined"
// when calling main twice while testing.
func init() {
	flag.BoolVar(&showVersion, "version", false, "print version and exit")
	flag.BoolVar(&quietProcess, "quiet", false, "hide process output")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if showVersion {
		fmt.Println("timer version", version)
	} else {
		processArgs(flag.Arg(0), flag.Arg(1))
	}
}

// If flag not specified in first position, it will be ignored.
// Example: ./timer start azerty -C ls
// --> flag '-C' will be ignored
func processArgs(cmd string, name string) {
	setUsage()
	checkParameters(cmd, name)
	executeTimerFunc(cmd, name)
}

func setUsage() {
	tpl, _ := template.New("usage").Parse(`{{range .}}  {{.CmdName}}: {{.CmdDesc}}
{{end}}
`)
	flag.Usage = func() {
		fmt.Fprintf(stderr, "Usage: %s [option] command name\n", os.Args[0])
		fmt.Fprintf(stderr, " COMMANDS:\n")
		tpl.Execute(stderr, cmdList)
		fmt.Fprintf(stderr, " OPTIONS:\n")
		flag.PrintDefaults()
		fmt.Fprintf(stderr, "\n")
		os.Exit(1)
	}
}

func checkParameters(cmd string, name string) {
	if flag.NArg() == 0 || flag.NArg() > 2 {
		flag.Usage()
	}
	if (cmd == "start" || cmd == "read" || cmd == "stop" || cmd == "exec") && name == "" {
		fmt.Fprintf(stderr, "Please specify name of timer\n")
		os.Exit(1)
	}
}

func executeTimerFunc(cmd string, name string) {
	i := indexOf(cmd)
	if i != -1 {
		cmdList[i].CmdFunc(timer, name)
	} else {
		flag.Usage()
	}
}

// check specified command against supported commands
// and return index if found, -1 otherwise
func indexOf(cmdName string) int {
	for i, item := range cmdList {
		if item.CmdName == cmdName {
			return i
		}
	}
	return -1
}
