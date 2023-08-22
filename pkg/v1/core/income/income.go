package coreincome

import (
	"github.com/11SF/inout-management/configs"
	"github.com/11SF/inout-management/pkg/v1/core/income/repository"
	"github.com/11SF/inout-management/pkg/v1/datamodel"
)

type Service interface {
	AddIncome(trans *datamodel.Transaction) error
	GetListIncome(uuid string) (transaction []*datamodel.Transaction, err error)
}

type service struct {
	db     repository.IIncomeRepositoryDB
	config *configs.Config
}

func NewService(db repository.IIncomeRepositoryDB, config *configs.Config) *service {
	return &service{
		db:     db,
		config: config,
	}
}
