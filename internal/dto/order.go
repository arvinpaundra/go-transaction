package dto

import "clean-arch/internal/model"

type (
	CreateOrderReq struct {
		UserId       int                    `json:"user_id" binding:"required"`
		Shipment     string                 `json:"shipment" binding:"required"`
		GrandTotal   float64                `json:"grand_total" binding:"required"`
		OrderDetails []CreateOrderDetailReq `json:"order_details" binding:"required"`
	}

	OrderRes struct {
		Id           int              `json:"id"`
		UserId       int              `json:"user_id"`
		Shipment     string           `json:"shipment"`
		Status       string           `json:"status"`
		GrandTotal   float64          `json:"grand_total"`
		OrderDetails []OrderDetailRes `json:"order_details"`
	}
)

func ToOrdersResponse(orders []model.Order) []OrderRes {
	var res []OrderRes

	for _, order := range orders {
		res = append(res, OrderRes{
			Id:           order.Id,
			UserId:       order.UserId,
			Shipment:     order.Shipment,
			Status:       order.Status,
			GrandTotal:   order.GrandTotal,
			OrderDetails: ToOrderDetailsResponse(order.OrderDetails),
		})
	}

	return res
}
