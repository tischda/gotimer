package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

const version string = "1.2.0"

var t Timer

var process string
var showVersion bool

type cliCmd struct {
	CmdName string
	CmdDesc string
	CmdFunc func(Timer, string)
}

var cmdList = []cliCmd{
	cliCmd{"start", "start timer", Timer.start},
	cliCmd{"read", "read timer (elapsed time)", Timer.read},
	cliCmd{"clear", "clear timer", Timer.clear},
	cliCmd{"stop", "read and then clear timer", Timer.stop},
	cliCmd{"list", "list timers", Timer.list},
}

func setUsage() {
	tmpl, _ := template.New("usage").Parse(`{{range .}}  {{.CmdName}}: {{.CmdDesc}}
{{end}}
`)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s command name\n       %s option\n\n", os.Args[0], os.Args[0])
		fmt.Fprintf(os.Stderr, " COMMANDS:\n")
		tmpl.Execute(os.Stderr, cmdList)
		fmt.Fprintf(os.Stderr, " OPTIONS:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(1)
	}
}

func main() {

	// no timestamp in logging
	log.SetFlags(0)

	flag.StringVar(&process, "C", "REQUIRED", "print elapsed time for process")
	flag.BoolVar(&showVersion, "version", false, "print version and exit")

	setUsage()

	flag.Parse()

	if showVersion {
		fmt.Println("timer version", version)
		return
	}

	processArgs(flag.Arg(0), flag.Arg(1))
}

func processArgs(cmd string, name string) {

	if flag.NArg() == 0 && flag.NFlag() != 1 {
		flag.Usage()
	}

	if cmd == "start" && name == "" {
		log.Fatalln("Please specify name of timer to start")
	}

	// execute process (http://bit.ly/1dMD2YN)
	if process != "REQUIRED" {
		t.process("cmd", "/c", process)
		return
	}

	i := commandIndex(cmd)
	if i == -1 {
		flag.Usage()
	}
	// execute command
	cmdList[i].CmdFunc(t, name)
}

func commandIndex(cmdName string) int {
	for i, item := range cmdList {
		if item.CmdName == cmdName {
			return i
		}
	}
	return -1
}
