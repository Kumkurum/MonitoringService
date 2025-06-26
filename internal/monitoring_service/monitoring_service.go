package monitoring_service

import (
	monitoringGrpc "MonitoringService/internal/transport"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"time"
)

type MonitoringService struct {
	//unknown string
}

// GetCpuInfo -Получение средней загрузки всех процессоров
func (s *MonitoringService) GetCpuInfo() (float64, error) {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return cpuPercent[0], err
}

// GetMemoryInfo - Информация о загрузке оперативной памяти
func (s *MonitoringService) GetMemoryInfo() (float64, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return memInfo.UsedPercent, err
}

// GetDiskInfo - Информация о загрузке жесткого диска
func (s *MonitoringService) GetDiskInfo() (float64, error) {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		fmt.Printf("Ошибка получения информации о загрузке диска: %v\n", err)
		return 0, err
	}
	return diskUsage.UsedPercent, err
}

// GetNetInfo - Информация о загрзке сети в срезе 1 секунды
func (s *MonitoringService) GetNetInfo() (uint64, uint64, error) {
	start, err0 := net.IOCounters(false)
	if err0 != nil {
		fmt.Printf("Ошибка получения информации о сети: %v\n", err0)
		return 0, 0, err0
	}
	time.Sleep(1 * time.Second)
	end, err1 := net.IOCounters(false)
	if err1 != nil {
		fmt.Printf("Ошибка получения информации о сети: %v\n", err1)
		return 0, 0, err1
	}
	return end[0].BytesSent - start[0].BytesSent, end[0].BytesRecv - start[0].BytesRecv, nil
}

// GetProcessesInfo - Информация об интересующих процессах если передать nil, то о всех процессах
func (s *MonitoringService) GetProcessesInfo(processName map[string]bool) ([]*monitoringGrpc.ProcessInfo, error) {
	processes, err := process.Processes()
	if err != nil {
		fmt.Printf("Ошибка получения процессов: %v\n", err)
		return nil, err
	}

	var processInfo = make([]*monitoringGrpc.ProcessInfo, 0)
	if processName == nil {
		fmt.Printf("ВСЕ!")
		for _, p := range processes {
			name, _ := p.Name()
			cpuPercent, _ := p.CPUPercent()
			memPercent, _ := p.MemoryPercent()
			status, _ := p.Status()
			processInfo = append(processInfo, &monitoringGrpc.ProcessInfo{
				Name:       name,
				CpuPercent: cpuPercent,
				MemPercent: memPercent,
				Status:     status[0],
			})
		}
		return processInfo, nil
	}
	for _, p := range processes {
		name, _ := p.Name()
		if processName[name] {
			cpuPercent, _ := p.CPUPercent()
			memPercent, _ := p.MemoryPercent()
			status, _ := p.Status()
			processInfo = append(processInfo, &monitoringGrpc.ProcessInfo{
				Name:       name,
				CpuPercent: cpuPercent,
				MemPercent: memPercent,
				Status:     status[0],
			})
		}
	}
	return processInfo, nil
}

// GetListenPort - Информация, какие порты находятся в состоянии LISTEN
func (s *MonitoringService) GetListenPort() ([]uint32, error) {
	// Получаем все сетевые соединения
	connections, err := net.Connections("all")
	if err != nil {
		fmt.Printf("Ошибка получения сетевых соединений: %v\n", err)
		return nil, err
	}
	var portInfo = make([]uint32, 0)
	for _, conn := range connections {
		if conn.Status == "LISTEN" {
			portInfo = append(portInfo, conn.Laddr.Port)
		}
	}
	return portInfo, nil
}
