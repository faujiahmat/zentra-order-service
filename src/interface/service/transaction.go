package service

import (
	"context"

	"github.com/faujiahmat/zentra-order-service/src/model/dto"
	"github.com/faujiahmat/zentra-order-service/src/model/entity"
)

type Transaction interface {
	Create(ctx context.Context, data *dto.TransactionReq) (*dto.TransactionRes, error)
	HandleNotif(ctx context.Context, data *entity.Transaction) error
}
