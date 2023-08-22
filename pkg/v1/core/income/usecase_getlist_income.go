package coreincome

import (
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	response "github.com/11SF/inout-management/utils"
	"golang.org/x/exp/slog"
)

type GetListIncomeFunc func(uuid string) ([]*datamodel.Transaction, error)

func (s *service) GetListIncome(uuid string) (transaction []*datamodel.Transaction, err error) {
	transaction, err = s.db.GetIncome(uuid)
	if err != nil {
		slog.Info("fail to get transaction list", "with", err.Error())
		return transaction, response.WriteError("R0500", err)
	}
	return transaction, nil
}
