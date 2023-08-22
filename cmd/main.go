package main

import (
	"fmt"
	"os"

	"github.com/11SF/inout-management/configs"
	routers "github.com/11SF/inout-management/pkg"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {

	slog.Info("Starting Server")
	// config := &configs.Config{
	// 	AppEnviroment: "development",
	// 	AppPort:       8080,
	// 	AppConstants: configs.AppConstants{
	// 		TranssactionIncomeType:  "income",
	// 		TranssactionExpenseType: "expense",
	// 	},
	// }

	config := configs.LoadEnv()

	server := routers.NewRouters(nil, nil, config).InitRouters()
	startServer(server, config)
}

func startServer(server *gin.Engine, config *configs.Config) {
	server.Run(fmt.Sprintf(":%v", config.AppPort))
}
