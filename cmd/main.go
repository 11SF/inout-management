package main

import (
	"fmt"
	"os"

	"github.com/11SF/inout-management/configs"
	routers "github.com/11SF/inout-management/pkg"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {

	slog.Info("[server] starting")
	config := configs.LoadEnv()

	slog.Info("[database] connecting to database")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Bangkok", config.Database.Host, config.Database.Username, config.Database.Password, config.Database.DatabaseName, config.Database.Port, config.Database.SSL)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		slog.Error("[database] fail to connect database", "with", err.Error())
		panic(0)
	}
	slog.Info("[database] connect to database success")

	server := routers.NewRouters(db, nil, config).InitRouters()
	startServer(server, config)
}

func startServer(server *gin.Engine, config *configs.Config) {
	server.Run(fmt.Sprintf(":%v", config.AppPort))
}
