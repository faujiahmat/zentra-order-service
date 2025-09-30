package repository

import (
	"context"

	"github.com/faujiahmat/zentra-order-service/src/model/entity"
)

type Product interface {
	FindByOrderId(ctx context.Context, orderId string) ([]*entity.ProductOrder, error)
}
