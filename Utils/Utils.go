package Utils

import (
	"fmt"

	"github.com/mitchellh/go-ps"

	"EvilDaisy/Configs"
)

// AccessIPList ...
var AccessIPList chan string

// CheckProcess ...
func CheckProcess() bool {
	if processes, err := ps.Processes(); err == nil {
		for _, x := range processes {
			if x.Executable() == "123" {
				return true
			}
		}
	} else {
		if Configs.IsDebug {
			fmt.Println(err.Error())
		}
	}

	return false
}
