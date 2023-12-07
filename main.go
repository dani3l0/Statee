package main

import (
	"fmt"
	"statee/syslib/disks"
	"time"
)

func main() {
	for {
		hmm, _ := disks.GetDisk("sda")
		fmt.Println(hmm)
		time.Sleep(time.Second)
	}
}
