package dto

import "clean-arch/internal/model"

type (
	CreateOrderDetailReq struct {
		ProductName string  `json:"product_name" binding:"required"`
		Qty         int     `json:"qty" binding:"required"`
		Price       float64 `json:"price" binding:"required"`
	}

	OrderDetailRes struct {
		Id          int     `json:"id"`
		ProductName string  `json:"product_name"`
		Qty         int     `json:"qty"`
		Price       float64 `json:"price"`
	}
)

func ToOrderDetailsResponse(orderDetails []model.OrderDetail) []OrderDetailRes {
	var res []OrderDetailRes

	for _, detail := range orderDetails {
		res = append(res, OrderDetailRes{
			Id:          detail.Id,
			ProductName: detail.ProductName,
			Qty:         detail.Qty,
			Price:       detail.Price,
		})
	}

	return res
}
