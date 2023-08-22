package apiincome

import (
	"net/http"

	coreincome "github.com/11SF/inout-management/pkg/v1/core/income"
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	"github.com/gin-gonic/gin"
)

type addIncomeHandler struct {
	addIncome coreincome.AddIncomeFunc
}

type addIncomeRequest struct {
	Amount float64
}

func NewAddIncomeHandler(addIncome coreincome.AddIncomeFunc) *addIncomeHandler {
	return &addIncomeHandler{addIncome: addIncome}
}

func (h *addIncomeHandler) AddIncome(c *gin.Context) {

	request := &addIncomeRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	trans := &datamodel.Transaction{
		Amount: request.Amount,
	}

	err = h.addIncome(trans)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, nil)
}
