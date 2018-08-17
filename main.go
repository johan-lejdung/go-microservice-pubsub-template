package main

import (
	"github.com/johan-lejdung/go-microservice-pubsub-template/bootstrap"
)

func main() {
	app := bootstrap.Service()

	app.PubSub.Init()

	app.PubSub.Consume()
}
