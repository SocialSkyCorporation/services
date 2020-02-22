package main

import (
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-micro/v2"
	"test/handler"
	"test/subscriber"

	test "test/proto/test"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.test"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	test.RegisterTestHandler(service.Server(), new(handler.Test))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.test", service.Server(), new(subscriber.Test))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
