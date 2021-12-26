package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":3030")
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Could not serve %v", err)
	}
	log.Println("server is listening on port: 3030")
}
