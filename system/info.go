package system

import (
	"github.com/shirou/gopsutil/mem"
)

// Report contains memory information.
type Report struct {
	Memory *mem.VirtualMemoryStat `json:"memory"`
}

// Info returns current status of the machine.
func Info() *Report {
	report := &Report{}

	if memory, err := mem.VirtualMemory(); err != nil {
		return report
	} else {
		report.Memory = memory
	}

	return report
}
