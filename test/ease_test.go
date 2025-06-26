package test

import (
	monitoringGrpc "MonitoringService/internal/transport"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestRequest(t *testing.T) {
	conn, err := grpc.NewClient(
		"127.0.0.1:50052", // Адрес сервера
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Создаем клиент (используя сгенерированный код)
	client := monitoringGrpc.NewMonitoringGrpcServiceClient(conn)

	response, err := client.GetAllInfo(
		context.Background(),
		&monitoringGrpc.AllInfoRequest{
			ProcessName: []string{"jetbrains"},
			Token:       "default",
		},
	)
	if err != nil {
		t.Errorf("fail to dial: %v", err)
	}

	println(response.GetSuccess().CpuPercent)
}
