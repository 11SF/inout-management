package apiincome

import (
	"errors"
	"net/http"

	coreincome "github.com/11SF/inout-management/pkg/v1/core/income"
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	response "github.com/11SF/inout-management/utils"
	"github.com/gin-gonic/gin"
)

type getListIncomeHandler struct {
	getListIncome coreincome.GetListIncomeFunc
}

type getListIncomeRequest struct {
	UUID string `json:"uuid"`
}

type getListIncomeResponse struct {
	IncomeList []*datamodel.Transaction `json:"income_list"`
}

func NewgetListIncomeHandler(getListIncome coreincome.GetListIncomeFunc) *getListIncomeHandler {
	return &getListIncomeHandler{getListIncome: getListIncome}
}

func (h *getListIncomeHandler) GetListIncome(c *gin.Context) {

	request := &getListIncomeRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(response.NewResponseError(http.StatusBadRequest, nil, response.WriteError("H0400", errors.New("invalid request"))))
		return
	}

	transactionList, err := h.getListIncome("") //mock uuid
	if err != nil {
		c.JSON(response.NewResponseError(http.StatusInternalServerError, nil, err))
		return
	}

	c.JSON(response.NewResponse(http.StatusOK, &getListIncomeResponse{IncomeList: transactionList}))
}
