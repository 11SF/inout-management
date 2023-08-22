package repository

import (
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	"gorm.io/gorm"
)

type IncomeRepositoryDB struct {
	db *gorm.DB
}

type IIncomeRepositoryDB interface {
	AddIncome(trans *datamodel.Transaction) error
}

func NewIncomeRepositoryDB(db *gorm.DB) *IncomeRepositoryDB {
	return &IncomeRepositoryDB{db}
}

func (repo *IncomeRepositoryDB) AddIncome(trans *datamodel.Transaction) error {
	if err := repo.db.Create(&trans).Error; err != nil {
		return err
	}
	return nil
}
