package main

import (
	"fmt"
	"net/http"
	"statee/api/machine"
)

func main() {
	http.HandleFunc("/api/cpu", machine.GetCpu)
	http.HandleFunc("/api/memory", machine.GetMemory)
	http.HandleFunc("/api/disks", machine.GetDisks)
	http.HandleFunc("/api/network", machine.GetNetwork)
	http.HandleFunc("/api/osinfo", machine.GetOsInfo)
	http.HandleFunc("/api/processes", machine.GetProcesses)

	listen_addr := ":9090"
	fmt.Println("Running on " + listen_addr)
	http.ListenAndServe(listen_addr, nil)
}
