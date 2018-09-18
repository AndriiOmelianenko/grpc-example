package api

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Receive message %s %s", in.Greeting, in.Name)
	return &PingMessage{
		Greeting: "Hello",
		Name:     in.Name,
	}, nil
}
