package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/markuscandido/go-expert-courses-crud/internal/application"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/adapter/config"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/adapter/grpc/v1/pb"
	"github.com/markuscandido/go-expert-courses-crud/internal/application/adapter/grpc/v1/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(cfg *config.Config, useCases *application.UseCases) {
	categoryService := service.NewCategoryService(
		useCases.CreateCategoryUseCase,
	)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("FATAL: failed to listen for gRPC: %v", err)
	}

	log.Printf("Starting gRPC server on port %s...", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("FATAL: could not start gRPC server: %v", err)
	}
}
