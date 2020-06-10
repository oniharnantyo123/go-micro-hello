package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	greeting "greeting/proto/greeting"
)

type Greeting struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Greeting) Call(ctx context.Context, req *greeting.Request, rsp *greeting.Response) error {
	log.Info("Received Greeting.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Greeting) Stream(ctx context.Context, req *greeting.StreamingRequest, stream greeting.Greeting_StreamStream) error {
	log.Infof("Received Greeting.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&greeting.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Greeting) PingPong(ctx context.Context, stream greeting.Greeting_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&greeting.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
