package repository

import (
	"github.com/11SF/inout-management/configs"
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	"gorm.io/gorm"
)

type ExpenseRepositoryDB struct {
	db     *gorm.DB
	config *configs.Config
}

type IExpenseRepositoryDB interface {
	AddExpense(trans *datamodel.Transaction) error
	GetExpense(uuid string) ([]*datamodel.Transaction, error)
}

func NewExpenseRepository(db *gorm.DB, config *configs.Config) *ExpenseRepositoryDB {
	db.AutoMigrate(&datamodel.Transaction{})
	return &ExpenseRepositoryDB{db, config}
}

func (repo *ExpenseRepositoryDB) AddExpense(trans *datamodel.Transaction) error {
	if err := repo.db.Create(&trans).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ExpenseRepositoryDB) GetExpense(uuid string) ([]*datamodel.Transaction, error) {
	transactions := []*datamodel.Transaction{}
	err := repo.db.Where("user_uuid = ? and transaction_type = ?", uuid, repo.config.AppConstants.TransactionExpenseType).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
