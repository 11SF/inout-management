package repository

import (
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	"gorm.io/gorm"
)

type ExpenseRepositoryDB struct {
	db *gorm.DB
}

type IExpenseRepositoryDB interface {
	AddExpense(trans *datamodel.Transaction) error
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepositoryDB {
	db.AutoMigrate(&datamodel.Transaction{})
	return &ExpenseRepositoryDB{db}
}

func (repo *ExpenseRepositoryDB) AddExpense(trans *datamodel.Transaction) error {
	if err := repo.db.Create(&trans).Error; err != nil {
		return err
	}
	return nil
}
