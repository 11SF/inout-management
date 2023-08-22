package coreincome

import (
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	"golang.org/x/exp/slog"
)

type AddIncomeFunc func(trans *datamodel.Transaction) error

func (s *service) AddIncome(trans *datamodel.Transaction) error {

	slog.Info("Start to add income")
	trans.TransactionType = string(s.config.AppConstants.TranssactionIncomeType)
	err := s.db.AddIncome(trans)
	if err != nil {
		slog.Error("add income error with: ", err.Error())
		return err
	}
	slog.Info("Complete to add income")

	return nil
}
