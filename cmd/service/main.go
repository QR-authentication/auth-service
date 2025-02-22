package main

import (
	"fmt"
	"log"
	"net"

	authproto "github.com/QR-authentication/auth-proto/auth-proto"
	"github.com/QR-authentication/auth-service/internal/config"
	"github.com/QR-authentication/auth-service/internal/infra"
	"github.com/QR-authentication/auth-service/internal/repository/postgres"
	"github.com/QR-authentication/auth-service/internal/service"
	metrics_lib "github.com/QR-authentication/metrics-lib"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.MustLoad()

	DBRepo := postgres.New(cfg)
	defer DBRepo.Close()

	metrics, err := metrics_lib.New(cfg.Metrics.Host, cfg.Metrics.Port, cfg.Service.Name, cfg.Platform.Env)
	if err != nil {
		log.Fatal("failed to create metrics object: ", err)
	}
	defer metrics.Disconnect()

	authService := service.New(DBRepo, cfg.Security.AuthSigningKey)
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			infra.MetricsInterceptor(metrics),
		),
	)

	authproto.RegisterAuthServiceServer(grpcServer, authService)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Service.Port))
	if err != nil {
		log.Fatalf("failed to start TCP listener: %v", err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to start gRPC listener: %v", err)
	}
}
