package delivery

import (
	"context"

	"github.com/faujiahmat/zentra-order-service/src/model/entity"
)

type ProductGrpc interface {
	ReduceStocks(ctx context.Context, data []*entity.ProductOrder) error
	RollbackStocks(ctx context.Context, data []*entity.ProductOrder) error
}
