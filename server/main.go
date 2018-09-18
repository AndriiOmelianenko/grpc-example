package main

import (
	"fmt"
	"log"
	"net"

	"github.com/AndriiOmelianenko/grpc-example/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// main starts a gRPC server and waits for connection
func main() {
	listenInterface := "localhost"
	listenPort := 7777

	// create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", listenInterface, listenPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Listening on %s:%d", listenInterface, listenPort)

	// create a server instance
	s := api.Server{}

	// Create the TLS credentials
	creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")
	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	// create a gRPC server object
	gRPCServer := grpc.NewServer(opts...)

	// attach the Ping service to the server
	api.RegisterPingServer(gRPCServer, &s)

	// start the server
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
