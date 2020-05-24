package app

import (
	"context"
	"fmt"
	"io"
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

func (a *app) EchoStream(s example.YourService_EchoStreamServer) error {

	messages := make([]string, 0)
	for {
		msg, err := s.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("ERROR", err)
		}

		messages = append(messages, msg.Value)
	}

	for i, msg := range messages {
		err := s.Send(&example.StringMessage{Value: fmt.Sprintf("num %d str = %s", i, msg)})
		if err != nil {
			log.Println("ERROR", err)
		}
	}

	return nil
}
