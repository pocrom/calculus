package calculus

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Add(ctx context.Context, message *Data) (*Result, error) {
	log.Printf("Received %d and %d ", message.GetA(), message.GetB())
	result := message.GetA() + message.GetB()

	return &Result{Result: result}, nil
}

func (s *Server) Sub(ctx context.Context, message *Data) (*Result, error) {
	result := message.A - message.B
	return &Result{Result: result}, nil
}
