package main

import (
	"fmt"
	"statee/syslib/osinfo"
	"time"
)

func main() {
	for {
		hmm := osinfo.GetOsInfo()
		fmt.Println(hmm)
		time.Sleep(time.Second)
	}
}
