package apiincome

import (
	"errors"
	"net/http"

	coreincome "github.com/11SF/inout-management/pkg/v1/core/income"
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	response "github.com/11SF/inout-management/utils"
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
		c.JSON(response.NewResponseError(http.StatusBadRequest, nil, response.WriteError("H0400", errors.New("invalid request"))))
		return
	}

	trans := &datamodel.Transaction{
		Amount: request.Amount,
	}

	err = h.addIncome(trans)
	if err != nil {
		c.JSON(response.NewResponseError(http.StatusInternalServerError, response.WriteError("U0500", err)))
		return
	}

	c.JSON(response.NewResponse(http.StatusOK, nil))
}
