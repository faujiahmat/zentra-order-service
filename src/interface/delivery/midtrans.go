package delivery

import (
	"context"

	"github.com/faujiahmat/zentra-order-service/src/model/dto"
)

type MidtransRESTful interface {
	Transaction(ctx context.Context, data *dto.TransactionReq) (*dto.MidtransTxRes, error)
}
