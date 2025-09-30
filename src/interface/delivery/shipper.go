package delivery

import (
	"context"

	"github.com/faujiahmat/zentra-order-service/src/model/entity"
)

type ShipperRESTful interface {
	ShippingOrder(ctx context.Context, data *entity.OrderWithProducts) (shippingId string, err error)
}
