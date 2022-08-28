package simple

import (
	"local.packages/share"
	"time"
)

type Deliverable interface {
	Deliver(deliveryAddress share.Address, deliveryDate time.Time) (DeliveringOrder, error)
}
