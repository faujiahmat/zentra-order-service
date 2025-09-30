package broker

import (
	"github.com/faujiahmat/zentra-order-service/src/core/broker/consumer"
	"github.com/faujiahmat/zentra-order-service/src/core/broker/handler"
	"github.com/faujiahmat/zentra-order-service/src/interface/service"
)

func InitKafkaConsumer(ts service.Transaction) *consumer.MidtransKafka {
	midtransHandler := handler.NewMidtransKafka(ts)
	midtransConsumer := consumer.NewMidtransKafka(midtransHandler)

	return midtransConsumer
}
