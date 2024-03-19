package main

import (
	"log"
	"strings"

	"github.com/mitchellh/go-ps"
)

func getProcessList() []ps.Process {
	psps, err := ps.Processes()
	if err != nil {
		panic(err)
	}

	return psps
}

// Search for a process by name and returns it if found.
func GetProcessByName(name string) (bool, ps.Process) {
	log.Printf("searching for %s ...", name)
	for _, ps := range getProcessList() {
		exName := ps.Executable()
		if strings.Contains(exName, name) {
			log.Printf("found! NAME: %v PID: %v PPID %v", exName, ps.Pid(), ps.PPid())
			return true, ps
		}
	}
	return false, nil
}
