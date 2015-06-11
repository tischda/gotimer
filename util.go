package main

import (
	"log"
	"regexp"
)

// splits a registry path to parent and child components
func splitPathSubkey(path regPath) (regPath, string) {
	regexp := regexp.MustCompile(`(.*)\\([^\\]+)$`)
	parts := regexp.FindStringSubmatch(path.lpSubKey)
	return regPath{path.hKeyIdx, parts[1]}, parts[2]
}

func checkFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
