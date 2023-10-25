package main

import (
	"fmt"
	"statee/syslib/cpu"
	"time"
)

func main() {
	for {
		hmm := cpu.Get()
		fmt.Println(hmm)
		time.Sleep(time.Second)
	}
}
