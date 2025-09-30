package client

import "github.com/faujiahmat/zentra-order-service/src/interface/delivery"

type Restful struct {
	Midtrans delivery.MidtransRESTful
	Shipper  delivery.ShipperRESTful
}

func NewRestful(md delivery.MidtransRESTful, sd delivery.ShipperRESTful) *Restful {
	return &Restful{
		Midtrans: md,
		Shipper:  sd,
	}
}
