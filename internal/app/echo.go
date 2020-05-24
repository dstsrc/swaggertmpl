package app

import (
	"context"
	"log"

	example "github.com/dstsrc/swaggertmpl/pkg/pb"
)

type app struct {
}

func NewApp() *app {
	return &app{}
}

func (a *app) Echo(ctx context.Context, s *example.StringMessage) (*example.StringMessage, error) {
	log.Println("echo method")
	return &example.StringMessage{Value: s.GetValue() + " echo"}, nil
}
