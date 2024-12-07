package main

import (
	"log"
	"net"
	"nicolascastro/go/grpc/database"
	"nicolascastro/go/grpc/server"
	"nicolascastro/go/grpc/studentpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", ":5060") // tcp es el protocolo de red que vamos a usar y 5060 es el puerto donde vamos a escuchar las peticiones
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	repo, err := database.NewPostgresRepository("postgres://postgres:password@localhost:5432/postgres?sslmode=disable")

	server := server.NewStudentServer(repo)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	s := grpc.NewServer()

	studentpb.RegisterStudentServiceServer(s, server)

	reflection.Register(s)

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
