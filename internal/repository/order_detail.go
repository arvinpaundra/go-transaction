package repository

import (
	"clean-arch/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	Insert(ctx *gin.Context, orderDetails []model.OrderDetail) error
	FindAll(ctx *gin.Context) ([]model.OrderDetail, error)
}

type orderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepository{db: db}
}

func (r *orderDetailRepository) Insert(ctx *gin.Context, orderDetails []model.OrderDetail) error {
	err := r.db.WithContext(ctx).Model(&model.OrderDetail{}).Create(orderDetails).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *orderDetailRepository) FindAll(ctx *gin.Context) ([]model.OrderDetail, error) {
	panic("not implemented")
}
