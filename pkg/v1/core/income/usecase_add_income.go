package coreincome

import (
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	response "github.com/11SF/inout-management/utils"
	"golang.org/x/exp/slog"
)

type AddIncomeFunc func(trans *datamodel.Transaction) error

func (s *service) AddIncome(trans *datamodel.Transaction) error {

	slog.Info("Start to add income")
	trans.TransactionType = string(s.config.AppConstants.TransactionIncomeType)

	slog.Info("add income to database")
	err := s.db.AddIncome(trans)
	if err != nil {
		slog.Error("add income to database error with: ", err.Error())
		return response.NewError("RP500", err.Error())
	}
	slog.Info("add income to database success")

	slog.Info("Complete to add income")
	return nil
}
