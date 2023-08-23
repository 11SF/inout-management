package coreexpense

import (
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	response "github.com/11SF/inout-management/utils"
	"golang.org/x/exp/slog"
)

type GetListExpenseFunc func(uuid string) ([]*datamodel.Transaction, error)

func (s *service) GetListExpense(uuid string) ([]*datamodel.Transaction, error) {
	transaction, err := s.db.GetExpense(uuid)
	if err != nil {
		slog.Info("fail to get transaction list", "with", err.Error())
		return transaction, response.NewError("RP500", err.Error())
	}
	return transaction, nil
}
