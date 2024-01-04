package machine

import (
	"fmt"
	"statee/machine/cpu"
	"statee/machine/disks"
	"statee/machine/memory"
	"statee/machine/network"
	"statee/machine/osinfo"
	"statee/machine/processes"
)

func Dump() {
	_cpu := cpu.GetCpu()
	fmt.Printf("-> %#v\n\n", _cpu)

	_mem := memory.GetMemory()
	fmt.Printf("-> %#v\n\n", _mem)

	_disks := disks.GetDisks()
	fmt.Printf("-> %#v\n\n", _disks)

	_net := network.GetNetwork()
	fmt.Printf("-> %#v\n\n", _net)

	_os := osinfo.GetOsInfo()
	fmt.Printf("-> %#v\n\n", _os)

	_proc := processes.GetProcesses()
	fmt.Printf("-> %#v\n\n", _proc)
}
