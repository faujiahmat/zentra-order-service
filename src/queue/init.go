package queue

import (
	restfulclient "github.com/faujiahmat/zentra-order-service/src/core/restful/client"
	"github.com/faujiahmat/zentra-order-service/src/interface/queue"
	"github.com/faujiahmat/zentra-order-service/src/interface/repository"
	"github.com/faujiahmat/zentra-order-service/src/queue/client"
	"github.com/faujiahmat/zentra-order-service/src/queue/handler"
	"github.com/faujiahmat/zentra-order-service/src/queue/server"
)

func InitServer(rc *restfulclient.Restful, or repository.Order) *server.Queue {
	orderHandler := handler.NewOrderQueue(rc, or)
	queueServer := server.NewQueue(orderHandler)

	return queueServer
}

func InitClient() queue.Client {
	queueClient := client.NewQueue()

	return queueClient
}
