package main

import (
	"statee/machine"
	"time"
)

func main() {
	for {
		machine.Dump()
		time.Sleep(time.Second)
	}
}
