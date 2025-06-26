package grpc_service

import (
	"MonitoringService/internal/monitoring_service"
	_ "MonitoringService/internal/monitoring_service"
	monitoringGrpc "MonitoringService/internal/transport"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"time"
)

type GrpcService struct {
	monitoringGrpc.MonitoringGrpcServiceServer
	core  monitoring_service.MonitoringService
	token string
}

func (g *GrpcService) GetAllInfo(ctx context.Context, request *monitoringGrpc.AllInfoRequest) (*monitoringGrpc.AllInfoResponse, error) {

	if request.GetToken() != g.token {
		return &monitoringGrpc.AllInfoResponse{
				Response: &monitoringGrpc.AllInfoResponse_Error{
					Error: &monitoringGrpc.Error{Code: monitoringGrpc.Error_UNKNOWN_TOKEN},
				},
			},
			fmt.Errorf("invalid token")
	}
	cpuPercent, _ := g.core.GetCpuInfo()
	fmt.Printf("CPU: %.2f%%\n", cpuPercent)

	// Память
	memInfo, _ := g.core.GetMemoryInfo()
	fmt.Printf("Memory: %.2f%% used\n", memInfo)

	// Диск
	diskUsage, _ := g.core.GetDiskInfo()
	fmt.Printf("Disk: %.2f%% used\n", diskUsage)

	send, recv, _ := g.core.GetNetInfo()
	fmt.Printf("Network: Sent %v bytes, Recv %v bytes\n", send, recv)

	names := make(map[string]bool)
	for _, name := range request.ProcessName {
		names[name] = true
	}
	process, _ := g.core.GetProcessesInfo(names)
	fmt.Printf("Process: %v \n", process)

	ports, _ := g.core.GetListenPort()
	fmt.Printf("Listen ports: %d \n", ports)

	return &monitoringGrpc.AllInfoResponse{
		Response: &monitoringGrpc.AllInfoResponse_Success{
			Success: &monitoringGrpc.AllInfoResponseSuccess{
				CpuPercent:  cpuPercent,
				MemPercent:  memInfo,
				DiskPercent: diskUsage,
				NetInfoSent: send,
				NetInfoRecv: recv,
				ListenPorts: ports,
				ProcessInfo: process,
				Time:        timestamppb.New(time.Now()),
			},
		},
	}, nil
}

func NewGrpcService(addr string, token string) *GrpcService {
	s := grpc.NewServer()
	grpcService := &GrpcService{token: token, core: monitoring_service.MonitoringService{}}

	monitoringGrpc.RegisterMonitoringGrpcServiceServer(s, grpcService)

	// Открыть порт 50051 для приема сообщений
	lis, err := net.Listen("tcp", ":"+addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Начать цикл приема и обработку запросов
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return grpcService
}
