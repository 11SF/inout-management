package coreexpense

import (
	"github.com/11SF/inout-management/configs"
	"github.com/11SF/inout-management/pkg/v1/core/expense/repository"
	"github.com/11SF/inout-management/pkg/v1/datamodel"
)

type Service interface {
	AddExpense(trans *datamodel.Transaction) error
}

type service struct {
	db     repository.IExpenseRepositoryDB
	config *configs.Config
}

func NewService(db repository.IExpenseRepositoryDB, config *configs.Config) *service {
	return &service{
		db:     db,
		config: config,
	}
}
