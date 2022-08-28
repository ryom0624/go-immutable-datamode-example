package simple

import (
	"local.packages/share"
	"log"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	inProgressOrder, err := NewInProgressOrder(1)
	if err != nil {
		log.Fatalln(err)
	}

	deliveringOrder, err := inProgressOrder.Deliver(
		share.Address{
			Country:       "JP",
			PostalCode:    "1510001",
			Region:        "Tokyo",
			StreetAddress: "Shibuya-ku",
		},
		time.Date(2022, 8, 29, 0, 0, 0, 0, time.Local),
	)
	if err != nil {
		log.Fatalln(err)
	}

	deliveringOrder.Close()

	inProgressOrder, err = NewInProgressOrder(2)
	if err != nil {
		log.Fatalln(err)
	}

	deliveringOrder, err = inProgressOrder.Deliver(
		share.Address{
			Country:       "US",
			PostalCode:    "1510001",
			Region:        "Tokyo",
			StreetAddress: "Shibuya-ku",
		},
		time.Date(2022, 8, 29, 0, 0, 0, 0, time.Local),
	)
	if err == nil {
		log.Fatalln(err)
	} else {
		return
	}

}
