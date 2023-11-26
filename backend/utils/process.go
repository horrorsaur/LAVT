package utils

import (
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

func searchProcessName(name string) bool {
	for _, ps := range getProcessList() {
		exName := ps.Executable()
		if strings.Contains(exName, name) {
			return true
		}
	}

	return false
}
