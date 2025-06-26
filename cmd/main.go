package main

import (
	"MonitoringService/internal/grpc_service"
	"flag"
	"fmt"
)

func main() {
	var addr, token string
	var version, help bool
	flag.StringVar(&addr, "g_addr", "50052", "address of grpc service")
	flag.StringVar(&token, "token", "default", "verification token for grpc service")
	flag.BoolVar(&version, "version", false, "Version service")
	flag.BoolVar(&help, "help", false, "Help how to use service")
	flag.Parse()
	if version {
		fmt.Println("Version 0.0.1")
		return
	}
	if help {
		fmt.Println("This is a service for Monitoring params in server machine")
		fmt.Println("flag addr to set port for grpc service, default = 50052 ")
		fmt.Println("flag token to set verification token, default = default")
		fmt.Println("just easy to use!")
		return
	}
	grpc_service.NewGrpcService(addr, token)
}
