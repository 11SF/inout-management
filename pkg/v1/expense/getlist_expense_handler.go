package apiexpense

import (
	"net/http"

	coreexpense "github.com/11SF/inout-management/pkg/v1/core/expense"
	"github.com/11SF/inout-management/pkg/v1/datamodel"
	response "github.com/11SF/inout-management/utils"
	"github.com/gin-gonic/gin"
)

type getListExpenseHandler struct {
	getListExpense coreexpense.GetListExpenseFunc
}

type getListExpenseRequest struct {
	UUID string `json:"uuid"`
}

type getListExpenseResponse struct {
	ExpenseList []*datamodel.Transaction `json:"Expense_list"`
}

func NewgetListExpenseHandler(getListExpense coreexpense.GetListExpenseFunc) *getListExpenseHandler {
	return &getListExpenseHandler{getListExpense: getListExpense}
}

func (h *getListExpenseHandler) GetListExpense(c *gin.Context) {

	request := &getListExpenseRequest{}
	err := c.BindJSON(&request)
	if err != nil {
		response.NewResponseError(c, http.StatusBadRequest, response.NewError("RQ400", "invalid request"))
		return
	}

	transactionList, err := h.getListExpense("") //mock uuid
	if err != nil {
		response.NewResponseError(c, http.StatusInternalServerError, err)
		return
	}

	response.NewResponse(c, http.StatusOK, &getListExpenseResponse{ExpenseList: transactionList})
}
