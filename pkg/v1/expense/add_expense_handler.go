package apiexpense

import (
	"net/http"

	coreexpense "github.com/11SF/inout-management/pkg/v1/core/expense"
	"github.com/11SF/inout-management/pkg/v1/datamodel"
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
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	trans := &datamodel.Transaction{
		Amount: request.Amount,
	}

	err = h.addExpense(trans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
