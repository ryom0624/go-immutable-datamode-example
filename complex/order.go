package main

import (
	"errors"
	"local.packages/share"
	"time"
)

type Order struct {
	orderId         int
	status          OrderStatus
	deliveryAddress share.Address
	deliveryDate    time.Time
}

func NewOrder(id int) *Order {
	return &Order{
		orderId: id,
	}
}

func (o *Order) Deliver(deliveryAddress share.Address, deliveryDate time.Time) error {
	if o.status != InProgress {
		return errors.New("order status is not in progress")
	}

	o.deliveryAddress = deliveryAddress
	o.deliveryDate = deliveryDate
	o.status = Delivering

	return o.Validate()
}

func (o Order) Validate() error {

	switch o.deliveryDate.Weekday() {
	case 1, 2, 3, 4, 5:
		return nil
	case 0, 6:
		return errors.New("This order can deliver only on weekday")
	}

	if o.deliveryAddress.Country != "JP" {
		return errors.New("This order is shipping only for Japan")
	}

	return nil
}
