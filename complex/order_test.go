package main

import (
	"local.packages/share"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	order := NewOrder(1)
	if err := order.Deliver(
		share.Address{
			Country:       "JP",
			PostalCode:    "1500001",
			Region:        "Tokyo",
			StreetAddress: "Shibuya-ku",
		},
		time.Date(2022, 8, 28, 0, 0, 0, 0, time.Local),
	); err == nil {
		t.Errorf("unexpected error")
	}

	order = NewOrder(2)
	if err := order.Deliver(
		share.Address{
			Country:       "JP",
			PostalCode:    "1500001",
			Region:        "Tokyo",
			StreetAddress: "Shibuya-ku",
		},
		time.Date(2022, 8, 29, 0, 0, 0, 0, time.Local),
	); err != nil {
		t.Errorf("unexpected error")
	}

	order = NewOrder(3)
	if err := order.Deliver(
		share.Address{
			Country:       "US",
			PostalCode:    "",
			Region:        "Seatle",
			StreetAddress: "",
		},
		time.Date(2022, 8, 29, 0, 0, 0, 0, time.Local),
	); err != nil {
		t.Errorf("unexpected error")
	}

}
