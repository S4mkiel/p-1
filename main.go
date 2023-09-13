package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	service "github.com/S4mkiel/p-1/domain/module"
	"github.com/S4mkiel/p-1/infra/db"
	config "github.com/S4mkiel/p-1/infra/env"
	"github.com/S4mkiel/p-1/infra/http"
	"github.com/S4mkiel/p-1/infra/logger"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	if os.Getenv("ENV") != "production" {
		LoadConfig()
	}
	fx.New(
		config.Module,
		service.Module,
		db.Module,
		logger.Module,
		http.Module,
	).Run()
}

func LoadConfig() {
	_, b, _, _ := runtime.Caller(0)

	basepath := filepath.Dir(b)

	err := godotenv.Load(fmt.Sprintf("%v/.env", basepath))
	if err != nil {
		panic(err)
	}
}
