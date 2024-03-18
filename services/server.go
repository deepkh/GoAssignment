package services

import (
	"fmt"

	"go-recommendation-system/protos"
	"log"
	"net"

	v1 "go-recommendation-system/services/v1"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func RunGrpcServer() {

	addr := fmt.Sprintf("0.0.0.0:5001")
	log.Printf("listening to %v", addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	services := v1.NewServices()
	gs := grpc.NewServer()
	protos.RegisterAuthServServer(gs, services)
	protos.RegisterRegServServer(gs, services)
	protos.RegisterRecommendationServServer(gs, services)
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
