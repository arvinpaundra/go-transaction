package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type TxBeginner interface {
	Begin(opts ...*sql.TxOptions) (Transaction, error)
}

type Transaction interface {
	Rollback() error
	Commit() error

	//
	UserRepository() UserRepository
	OrderRepository() OrderRepository
	OrderDetailRepository() OrderDetailRepository
}

type txBeginner struct {
	db *gorm.DB
}

func NewTxBeginner(db *gorm.DB) TxBeginner {
	return &txBeginner{db: db}
}

func (t *txBeginner) Begin(opts ...*sql.TxOptions) (Transaction, error) {
	tx := t.db.Begin(opts...)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &transaction{db: tx}, nil
}

type transaction struct {
	db *gorm.DB
}

func (tx *transaction) Rollback() error {
	err := tx.db.Rollback().Error
	if err != nil {
		return err
	}

	return nil
}

func (tx *transaction) Commit() error {
	err := tx.db.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (tx *transaction) UserRepository() UserRepository {
	return NewUserRepository(tx.db)
}

func (tx *transaction) OrderRepository() OrderRepository {
	return NewOrderRepository(tx.db)
}

func (tx *transaction) OrderDetailRepository() OrderDetailRepository {
	return NewOrderDetailRepository(tx.db)
}
