package apiexpense

import (
	"errors"
	"net/http"

	coreexpense "github.com/11SF/inout-management/pkg/v1/core/expense"
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	response "github.com/11SF/inout-management/utils"
	"github.com/gin-gonic/gin"
)

type addExpenseHandler struct {
	addExpense coreexpense.AddExpenseFunc
}

type addExpenseRequest struct {
	Amount float64
}

func NewAddExpenseHandler(addExpense coreexpense.AddExpenseFunc) *addExpenseHandler {
	return &addExpenseHandler{addExpense: addExpense}
}

func (h *addExpenseHandler) AddExpense(c *gin.Context) {

	request := &addExpenseRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		response.NewResponseError(http.StatusBadRequest, nil, response.WriteError("H0400", errors.New("invalid request")))
		return
	}

	trans := &datamodel.Transaction{
		Amount: request.Amount,
	}

	err = h.addExpense(trans)
	if err != nil {
		response.NewResponseError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
