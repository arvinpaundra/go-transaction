package repository

import (
	"clean-arch/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Insert(ctx *gin.Context, order *model.Order) error
	FindAll(ctx *gin.Context) ([]model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Insert(ctx *gin.Context, order *model.Order) error {
	err := r.db.WithContext(ctx).Model(&model.Order{}).Create(order).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *orderRepository) FindAll(ctx *gin.Context) ([]model.Order, error) {
	panic("not implemented")
}
