package order

import (
	"clean-arch/internal/dto"
	"clean-arch/internal/factory"
	"clean-arch/internal/model"
	"clean-arch/internal/repository"
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Insert(ctx *gin.Context, req *dto.CreateOrderReq) error
	FindAll(ctx *gin.Context) ([]dto.OrderRes, error)
}

type service struct {
	txBeginner            repository.TxBeginner
	userRepository        repository.UserRepository
	orderRepository       repository.OrderRepository
	orderDetailRepository repository.OrderDetailRepository
}

func NewService(f *factory.Factory) Service {
	return &service{
		txBeginner:            f.TxBeginner,
		userRepository:        f.UserRepository,
		orderRepository:       f.OrderRepository,
		orderDetailRepository: f.OrderDetailRepository,
	}
}

func (s *service) Insert(ctx *gin.Context, req *dto.CreateOrderReq) error {
	tx, err := s.txBeginner.Begin(&sql.TxOptions{})
	if err != nil {
		return err
	}

	// check user is exist
	_, err = tx.UserRepository().FindById(ctx, req.UserId)
	if err != nil {
		return err
	}

	// insert to order
	order := model.Order{
		UserId:     req.UserId,
		Shipment:   strings.ToUpper(req.Shipment),
		Status:     "created",
		GrandTotal: req.GrandTotal,
	}

	err = tx.OrderRepository().Insert(ctx, &order)
	if err != nil {
		tx.Rollback()
		return err
	}

	// insert the order details
	var orderDetails []model.OrderDetail

	for _, orderDetail := range req.OrderDetails {
		orderDetails = append(orderDetails, model.OrderDetail{
			OrderId:     order.Id,
			ProductName: orderDetail.ProductName,
			Qty:         orderDetail.Qty,
			Price:       orderDetail.Price,
		})
	}

	err = tx.OrderDetailRepository().Insert(ctx, orderDetails)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (s *service) FindAll(ctx *gin.Context) ([]dto.OrderRes, error) {
	panic("not implemented")
}
