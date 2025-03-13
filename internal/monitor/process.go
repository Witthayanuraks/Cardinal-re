package monitor

import (
	"github.com/shirou/gopsutil/process"
)

// ProcessInfo menyimpan info proses yang berjalan
type ProcessInfo struct {
	PID  int32
	Name string
	CPU  float64
}

// GetRunningProcesses mendapatkan daftar proses aktif
func GetRunningProcesses() []ProcessInfo {
	processes, _ := process.Processes()
	var result []ProcessInfo

	for _, p := range processes {
		name, _ := p.Name()
		cpuPercent, _ := p.CPUPercent()

		result = append(result, ProcessInfo{
			PID:  p.Pid,
			Name: name,
			CPU:  cpuPercent,
		})
	}
	return result
}
