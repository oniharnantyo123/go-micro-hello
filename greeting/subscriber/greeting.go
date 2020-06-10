package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	greeting "greeting/proto/greeting"
)

type Greeting struct{}

func (e *Greeting) Handle(ctx context.Context, msg *greeting.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *greeting.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
