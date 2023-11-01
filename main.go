package main

import (
	"fmt"
	"statee/syslib/network"
	"time"
)

func main() {
	for {
		hmm := network.GetInterfaces()
		fmt.Println(hmm)
		time.Sleep(time.Second)
	}
}
