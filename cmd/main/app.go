package main

import (
	"context"
	"github.com/dkischenko/xm_app/internal/app"
	"github.com/dkischenko/xm_app/internal/company"
	"github.com/dkischenko/xm_app/internal/company/database"
	"github.com/dkischenko/xm_app/internal/config"
	"github.com/dkischenko/xm_app/pkg/database/postgres"
	"github.com/dkischenko/xm_app/pkg/logger"
	"github.com/gorilla/mux"
	"os"
	"sync"
	"time"
)

func main() {
	l, err := logger.GetLogger()
	if err != nil {
		panic(err)
	}

	l.Entry.Info("Create router")
	router := mux.NewRouter()

	var cfg *config.Config
	var once sync.Once
	configPath := os.Getenv("CONFIG")
	once.Do(func() {
		cfg = config.GetConfig(configPath, &config.Config{})
	})

	l.Entry.Info("Create database connection")
	client, err := postgres.NewClient(context.Background(), cfg.Storage.Host, cfg.Storage.Port,
		cfg.Storage.Username, cfg.Storage.Password, cfg.Storage.Database)

	if err != nil {
		panic(err)
	}

	storage := database.NewStorage(client, l)
	accessTokenTTL, err := time.ParseDuration(cfg.Auth.AccessTokenTTL)
	if err != nil {
		panic(err)
	}

	service := company.NewService(l, storage, accessTokenTTL)
	handler := company.NewHandler(l, service, cfg)
	handler.Register(router)
	app.Run(router, l, cfg)
}
