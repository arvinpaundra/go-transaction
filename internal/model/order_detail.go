package model

type OrderDetail struct {
	Id          int
	OrderId     int
	ProductName string
	Qty         int
	Price       float64
	Common
}
