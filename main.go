package main

import (
	"statee/syslib/disks"
)

func main() {
	disks.GetSmartData("/dev/sda")
}
