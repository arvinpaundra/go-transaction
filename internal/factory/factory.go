package factory

import (
	"clean-arch/database"
	"clean-arch/internal/repository"
)

type Factory struct {
	TxBeginner            repository.TxBeginner
	UserRepository        repository.UserRepository
	OrderRepository       repository.OrderRepository
	OrderDetailRepository repository.OrderDetailRepository
}

func NewFactory() *Factory {
	// Check db connection
	db := database.GetConnection()
	return &Factory{
		// Pass the db connection to repository package for database query calling
		repository.NewTxBeginner(db),
		repository.NewUserRepository(db),
		repository.NewOrderRepository(db),
		repository.NewOrderDetailRepository(db),
	}
}
