package configs

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

type Config struct {
	AppEnviroment string `env:"APP_ENVIROMENT,required"`
	AppPort       int    `env:"APP_PORT,required"`
	AppConstants  AppConstants
	Database      Database
}

type AppConstant string

type AppConstants struct {
	TransactionIncomeType  AppConstant
	TransactionExpenseType AppConstant
}

type Database struct {
	Host         string `env:"DB_HOST,required"`
	Port         int    `env:"DB_PORT,required"`
	Username     string `env:"DB_USERNAME,required"`
	Password     string `env:"DB_PASSWORD,required"`
	DatabaseName string `env:"DB_DATABASE_NAME,required"`
	SSL          string `env:"DB_SSL,required"`
}

func LoadEnv() *Config {
	slog.Info("[env] start loading env")
	err := godotenv.Load("./configs/.env")
	if err != nil {
		slog.Error("[env] unable to load .env file", "error", err)
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		slog.Error("[env] unable to parse ennvironment variables", "error", err)
		panic(0)
	}

	cfg.AppConstants.TransactionIncomeType = "INCOME"
	cfg.AppConstants.TransactionExpenseType = "EXPENSE"

	slog.Info("[env] loading env complete")

	return cfg
}
