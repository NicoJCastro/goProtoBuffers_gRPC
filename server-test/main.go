package main

import (
	"log"
	"net"
	"nicolascastro/go/grpc/database"
	"nicolascastro/go/grpc/repository"
	"nicolascastro/go/grpc/server"
	"nicolascastro/go/grpc/testpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5070")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:password@host.docker.internal:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database")

	repository.SetRepository(repo)

	server := server.NewTestServer(repo)

	s := grpc.NewServer()

	testpb.RegisterTestServiceServer(s, server)

	reflection.Register(s)

	log.Println("Starting gRPC server on port 5070")

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
