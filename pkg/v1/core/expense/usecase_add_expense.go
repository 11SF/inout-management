package coreexpense

import (
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	response "github.com/11SF/inout-management/utils"
	"golang.org/x/exp/slog"
)

type AddExpenseFunc func(trans *datamodel.Transaction) error

func (s *service) AddExpense(trans *datamodel.Transaction) error {

	slog.Info("Start to add expense")
	trans.TransactionType = string(s.config.AppConstants.TransactionExpenseType)

	slog.Info("add expense to database")
	err := s.db.AddExpense(trans)
	if err != nil {
		slog.Error("add expense to database error with: ", err.Error())
		return response.NewError("RP500", err.Error())
	}
	slog.Info("add expense to database success")

	slog.Info("Complete to add expense")
	return nil
}
