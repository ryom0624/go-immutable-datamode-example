package simple

import (
	"errors"
	"local.packages/share"
	"time"
)

type DeliveringOrder struct {
	orderId         int
	DeliveryAddress share.Address
	DeliveryDate    time.Time
}

func NewDeliveringOrder(
	orderId int,
	deliveryAddress share.Address,
	deliveryDate time.Time,
) (*DeliveringOrder, error) {
	switch deliveryDate.Weekday() {
	case 1, 2, 3, 4, 5:
	case 0, 6:
		return nil, errors.New("This order can deliver only on weekday")
	}
	if deliveryAddress.Country != "JP" {
		return nil, errors.New("This order is shipping only for Japan")
	}

	return &DeliveringOrder{
		orderId:         orderId,
		DeliveryAddress: deliveryAddress,
		DeliveryDate:    deliveryDate,
	}, nil
}

func (o DeliveringOrder) GetOrderId() int {
	return o.orderId
}

func (o DeliveringOrder) Close() ClosedOrder {
	return ClosedOrder{}
}
