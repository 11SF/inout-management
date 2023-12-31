package apiincome

import (
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
	Amount   float64 `json:"amount"`
	UserUUID string  `json:"user_uuid"`
	Message  string  `json:"message"`
}

func NewAddIncomeHandler(addIncome coreincome.AddIncomeFunc) *addIncomeHandler {
	return &addIncomeHandler{addIncome: addIncome}
}

func (h *addIncomeHandler) AddIncome(c *gin.Context) {

	request := &addIncomeRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		response.NewResponseError(c, http.StatusBadRequest, response.NewError("RQ400", "invalid request"))
		return
	}

	trans := &datamodel.Transaction{
		Amount:   request.Amount,
		UserUUID: request.UserUUID,
		Message:  request.Message,
	}

	err = h.addIncome(trans)
	if err != nil {
		response.NewResponseError(c, http.StatusInternalServerError, err)
		return
	}

	response.NewResponse(c, http.StatusOK, nil)
}
