package routers

import (
	"github.com/11SF/inout-management/configs"
	coreexpense "github.com/11SF/inout-management/pkg/v1/core/expense"
	expenserepository "github.com/11SF/inout-management/pkg/v1/core/expense/repository"
	coreincome "github.com/11SF/inout-management/pkg/v1/core/income"
	incomerepository "github.com/11SF/inout-management/pkg/v1/core/income/repository"
	apiexpense "github.com/11SF/inout-management/pkg/v1/expense"
	apiincome "github.com/11SF/inout-management/pkg/v1/income"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Routers struct {
	db     *gorm.DB
	redis  *redis.Client
	config *configs.Config
}

func NewRouters(db *gorm.DB, redis *redis.Client, config *configs.Config) *Routers {
	return &Routers{db, redis, config}
}

func (router *Routers) InitRouters() *gin.Engine {
	if router.config.AppEnviroment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	contextPath := r.Group("/inout-management")
	v1 := contextPath.Group("/v1")

	expenseRepositoryDB := expenserepository.NewExpenseRepository(router.db, router.config)
	expenseService := coreexpense.NewService(expenseRepositoryDB, router.config)
	addExpenseHandler := apiexpense.NewAddExpenseHandler(expenseService.AddExpense)
	// getExpenseHandler := apiexpense
	v1.POST("/add-expense", addExpenseHandler.AddExpense)

	incomeRepositoryDB := incomerepository.NewIncomeRepositoryDB(router.db, router.config)
	incomeService := coreincome.NewService(incomeRepositoryDB, router.config)
	addIncomeHandler := apiincome.NewAddIncomeHandler(incomeService.AddIncome)
	getListIncomeHandler := apiincome.NewgetListIncomeHandler(incomeService.GetListIncome)
	v1.POST("/add-income", addIncomeHandler.AddIncome)
	v1.POST("/get-income", getListIncomeHandler.GetListIncome)

	return r
}
