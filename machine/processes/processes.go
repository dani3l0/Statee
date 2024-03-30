package processes

import (
	"os"
	"path"
	"statee/machine/utils"
	"strconv"
	"strings"
)

type Processes struct {
	Processes []Process
}

type Process struct {
	Name   string
	Memory int
	Pid    int
	Uid    int
}

// Get information of a process with provided id
func GetProcess(id int) (Process, error) {
	var process Process
	var err error
	raw, _ := utils.Cat(path.Join("/proc", strconv.Itoa(id), "status"))
	process.Name, err = utils.Grep(raw, "Name")
	if err != nil {
		return process, err
	}
	process.Name = strings.Trim(process.Name, "\t")
	process.Memory, err = utils.GrepInt(raw, "RssAnon")
	if err != nil {
		return process, err
	}
	process.Pid = id
	uid_raw, _ := utils.Grep(raw, "Uid")
	uid_raw = strings.Fields(uid_raw)[0]
	process.Uid, err = strconv.Atoi(uid_raw)
	return process, err
}

// List all PIDs
func GetProcesses() Processes {
	var processes Processes
	proc, _ := os.ReadDir("/proc")
	for _, process_id := range proc {
		pid, err := strconv.Atoi(process_id.Name())
		if err != nil {
			continue
		}
		process, err := GetProcess(pid)
		if err == nil {
			processes.Processes = append(processes.Processes, process)
		}
	}
	return processes
}
