package simple

import (
	"local.packages/share"
	"time"
)

type InProgressOrder struct {
	orderId int
}

func NewInProgressOrder(id int) (*InProgressOrder, error) {
	return &InProgressOrder{id}, nil
}

func (o InProgressOrder) GetOrderId() int {
	return o.orderId
}

func (o InProgressOrder) Deliver(
	deliveryAddress share.Address,
	deliveryDate time.Time,
) (*DeliveringOrder, error) {
	return NewDeliveringOrder(o.GetOrderId(), deliveryAddress, deliveryDate)
}
