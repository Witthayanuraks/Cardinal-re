package monitor

import (
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// SystemStats menyimpan info CPU, RAM, dll.
type SystemStats struct {
	CPUUsage    float64
	MemoryUsage uint64
	GoRoutines  int
}

// GetSystemStats mengambil informasi sistem
func GetSystemStats() SystemStats {
	cpuUsage, _ := cpu.Percent(0, false)
	memInfo, _ := mem.VirtualMemory()

	return SystemStats{
		CPUUsage:    cpuUsage[0],
		MemoryUsage: memInfo.Used / 1024 / 1024, // Convert ke MB
		GoRoutines:  runtime.NumGoroutine(),
	}
}
