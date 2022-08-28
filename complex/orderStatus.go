package main

type OrderStatus int

const (
	// InProgress - 商品手配中
	InProgress OrderStatus = iota

	// Delivering - 配送処理中
	Delivering

	// Closed - 出荷済み
	Closed
)
