package order_test

import (
	"clean-arch/internal/app/order"
	"clean-arch/internal/dto"
	"clean-arch/internal/factory"
	"clean-arch/internal/model"
	"clean-arch/internal/repository/mocks"
	"database/sql"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestTable struct {
	name string
	fn   func(t *testing.T)
}

var (
	txBeginner            mocks.TxBeginner
	transaction           mocks.Transaction
	userRepository        mocks.UserRepository
	orderRepository       mocks.OrderRepository
	orderDetailRepository mocks.OrderDetailRepository

	service order.Service

	userModel        model.User
	orderModel       model.Order
	orderDetailModel model.OrderDetail

	orderReq        dto.CreateOrderReq
	orderDetailsReq []dto.CreateOrderDetailReq

	ctx *gin.Context
)

func initService() {
	f := factory.Factory{
		TxBeginner:            &txBeginner,
		UserRepository:        &userRepository,
		OrderRepository:       &orderRepository,
		OrderDetailRepository: &orderDetailRepository,
	}

	service = order.NewService(&f)

	userModel = model.User{
		ID:    1,
		Name:  "test",
		Email: "test@mail.com",
	}

	orderModel = model.Order{
		Id:         1,
		UserId:     1,
		Shipment:   "JNE",
		Status:     "created",
		GrandTotal: 100000,
	}

	orderDetailModel = model.OrderDetail{
		Id:          1,
		OrderId:     orderModel.Id,
		ProductName: "test",
		Qty:         10,
		Price:       10000,
	}

	orderDetailsReq = []dto.CreateOrderDetailReq{
		{
			ProductName: orderDetailModel.ProductName,
			Qty:         orderDetailModel.Qty,
			Price:       orderDetailModel.Price,
		},
	}

	orderReq = dto.CreateOrderReq{
		UserId:       orderModel.UserId,
		Shipment:     orderModel.Shipment,
		GrandTotal:   orderModel.GrandTotal,
		OrderDetails: orderDetailsReq,
	}

	ctx = &gin.Context{}
}

func TestMain(m *testing.M) {

	initService()

	m.Run()
}

func TestInsert(t *testing.T) {
	tests := []TestTable{
		{
			name: "success",
			fn: func(t *testing.T) {
				txBeginner.On("Begin", &sql.TxOptions{}).Return(&transaction, nil).Once()

				transaction.On("UserRepository").Return(&userRepository).Once()

				userRepository.On("FindById", ctx, userModel.ID).Return(&userModel, nil).Once()

				transaction.On("OrderRepository").Return(&orderRepository).Once()

				orderRepository.On("Insert", ctx, mock.Anything).Return(nil).Once()

				transaction.On("OrderDetailRepository").Return(&orderDetailRepository).Once()

				orderDetailRepository.On("Insert", ctx, mock.Anything).Return(nil).Once()

				transaction.On("Commit").Return(nil).Once()

				err := service.Insert(ctx, &orderReq)

				assert.Nil(t, err)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.fn)
	}
}
