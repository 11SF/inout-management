package coreexpense

import (
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	"golang.org/x/exp/slog"
)

type AddExpenseFunc func(trans *datamodel.Transaction) error

func (s *service) AddExpense(trans *datamodel.Transaction) error {

	slog.Info("Start to add expense")
	trans.TransactionType = string(s.config.AppConstants.TranssactionExpenseType)
	err := s.db.AddExpense(trans)
	if err != nil {
		slog.Error("add expense error with: ", err.Error())
		return err
	}
	slog.Info("Complete to add expense")

	return nil
}
