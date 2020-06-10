package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"greeting/handler"
	"greeting/subscriber"

	greeting "greeting/proto/greeting"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.greeting"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	greeting.RegisterGreetingHandler(service.Server(), new(handler.Greeting))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.greeting", service.Server(), new(subscriber.Greeting))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
