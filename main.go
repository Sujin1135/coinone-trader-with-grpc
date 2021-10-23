package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO: register some interceptors
	s := grpc.NewServer()

	// TODO: register some proto objects

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
