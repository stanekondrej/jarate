package jarate

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type cpuStats struct {
	PerCpu []float64 `json:"per_cpu"`
	//Freq   uint      `json:"freq"`
}

func getCpuStats() (cpuStats, error) {
	perCpu, err := cpu.Percent(0, false)
	if err != nil {
		return cpuStats{}, err
	}

	// WARN: in the package, this is just hard-coded to 100. wtf.
	//freq := cpu.ClocksPerSec / 1_000_000

	return cpuStats{
		perCpu,
		//uint(freq),
	}, nil
}

// All of these are in bytes
type memStats struct {
	Used  uint64 `json:"used"`
	Total uint64 `json:"total"`
}

func getMemStats() (memStats, error) {
	vmem, err := mem.VirtualMemory()
	if err != nil {
		return memStats{}, err
	}

	return memStats{
		vmem.Used,
		vmem.Total,
	}, nil
}

type Stats struct {
	Cpu cpuStats `json:"cpu"`
	Mem memStats `json:"mem"`
}

func getStats() (Stats, error) {
	cpuStats, err := getCpuStats()
	if err != nil {
		return Stats{}, err
	}

	memStats, err := getMemStats()
	if err != nil {
		return Stats{}, err
	}

	return Stats{
		cpuStats,
		memStats,
	}, nil
}
