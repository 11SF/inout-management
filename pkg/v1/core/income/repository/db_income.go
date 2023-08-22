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
	GetIncome(uuid string) ([]*datamodel.Transaction, error)
}

func NewIncomeRepositoryDB(db *gorm.DB) *IncomeRepositoryDB {
	db.AutoMigrate(&datamodel.Transaction{})
	return &IncomeRepositoryDB{db}
}

func (repo *IncomeRepositoryDB) AddIncome(trans *datamodel.Transaction) error {
	if err := repo.db.Create(&trans).Error; err != nil {
		return err
	}
	return nil
}

func (repo *IncomeRepositoryDB) GetIncome(uuid string) ([]*datamodel.Transaction, error) {
	transactions := []*datamodel.Transaction{}
	err := repo.db.Where("user_uuid = ?", uuid).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
