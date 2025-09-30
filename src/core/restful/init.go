package restful

import (
	"github.com/faujiahmat/zentra-order-service/src/core/restful/client"
	"github.com/faujiahmat/zentra-order-service/src/core/restful/delivery"
	"github.com/faujiahmat/zentra-order-service/src/core/restful/handler"
	"github.com/faujiahmat/zentra-order-service/src/core/restful/middleware"
	"github.com/faujiahmat/zentra-order-service/src/core/restful/server"
	"github.com/faujiahmat/zentra-order-service/src/interface/service"
)

func InitServer(ts service.Transaction, os service.Order) *server.Restful {
	orderHandler := handler.NewOrderRESTful(ts, os)
	middleware := middleware.New()

	restfulServer := server.NewRestful(orderHandler, middleware)
	return restfulServer
}

func InitClient() *client.Restful {
	midtransDelivery := delivery.NewMidtransRESTful()
	shipperDelivery := delivery.NewShipperRESTful()
	restfulClient := client.NewRestful(midtransDelivery, shipperDelivery)

	return restfulClient
}
