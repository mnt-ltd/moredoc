package device

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type CPUInfo struct {
	Cores     int32   `json:"cores"`
	ModelName string  `json:"model_name"`
	Mhz       float64 `json:"mhz"`
	Percent   float64 `json:"percent"`
}

type MemInfo struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	Free        uint64  `json:"free"`
}

type DiskInfo struct {
	Total    uint64  `json:"total"`
	Used     uint64  `json:"used"`
	Free     uint64  `json:"free"`
	Percent  float64 `json:"percent"`
	DiskName string  `json:"disk_name"`
}

// GetMemory returns memory information
func GetMemory() (memInfo MemInfo) {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("mem.VirtualMemory() failed with %s", err)
		return
	}

	memInfo.Total = v.Total
	memInfo.Available = v.Available
	memInfo.Used = v.Used
	memInfo.UsedPercent = v.UsedPercent
	memInfo.Free = v.Free
	return
}

func GetDisk() (diskInfos []DiskInfo) {
	stats, _ := disk.Partitions(true)
	if runtime.GOOS != "windows" {
		usage, _ := disk.Usage("/")
		diskInfos = append(diskInfos, DiskInfo{
			DiskName: "/",
			Total:    usage.Total,
			Used:     usage.Used,
			Free:     usage.Free,
			Percent:  usage.UsedPercent,
		})
		return
	}

	for _, stat := range stats {
		usage, _ := disk.Usage(stat.Mountpoint)
		if usage != nil {
			diskInfos = append(diskInfos, DiskInfo{
				DiskName: stat.Mountpoint,
				Total:    usage.Total,
				Used:     usage.Used,
				Free:     usage.Free,
				Percent:  usage.UsedPercent,
			})
		}
	}
	return
}

// GetCPU returns CPU information
func GetCPU() (cpuInfo CPUInfo) {
	cpuInfos, _ := cpu.Info()
	for _, c := range cpuInfos {
		cpuInfo.Cores = c.Cores
		cpuInfo.ModelName = c.ModelName
		cpuInfo.Mhz = c.Mhz
	}

	percents, _ := cpu.Percent(time.Second, false)
	for _, p := range percents {
		cpuInfo.Percent = p
	}
	return
}
