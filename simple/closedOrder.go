package simple

import (
	"local.packages/share"
	"time"
)

type ClosedOrder struct {
	orderId         int
	DeliveryAddress share.Address
	DeliveryDate    time.Time
}

func (c ClosedOrder) GetOrderId() int {
	return c.orderId
}

func (c ClosedOrder) Validate() error {
	return nil
}
