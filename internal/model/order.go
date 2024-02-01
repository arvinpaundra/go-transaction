package model

type Order struct {
	Id           int
	UserId       int
	Shipment     string
	Status       string
	GrandTotal   float64
	OrderDetails []OrderDetail
	Common
}
