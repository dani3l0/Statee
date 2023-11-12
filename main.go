package main

import (
	"fmt"
	"statee/syslib/processes"
	"time"
)

func main() {
	for {
		hmm := processes.GetAllProcesses()
		fmt.Println(hmm)
		time.Sleep(time.Second)
	}
}
