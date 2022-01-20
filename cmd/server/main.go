package main

import (
	"github.com/dchebakov/tracker/config"
	"github.com/dchebakov/tracker/internal/server"
	"github.com/dchebakov/tracker/pkg/logger"
	"github.com/dchebakov/tracker/pkg/postgres"
	"github.com/go-playground/validator/v10"
)

// @title Tracker
// @version 0.1.0
// @description Track users statistics of valid/invalid API calls
// @host localhost:6000
// @BasePath /
func main() {
	cfg := config.New()
	logger := logger.NewLogger(cfg)

	logger.Info("Starting Tracker API")
	logger.Infof("LogLevel: %s, Mode: %s", cfg.Api.LogLevel, cfg.Api.Mode)

	logger.Info("Connecting to database...")
	db, err := postgres.NewDB(cfg)
	if err != nil {
		logger.Fatalf("Failed to connect to database", err)
	} else {
		logger.Infof("Connected to database, status: %#v", db.Stats())
	}

	validate := validator.New()
	s := server.NewServer(cfg, db, logger, validate)
	err = s.Run()
	if err != nil {
		logger.Fatal(err)
	}
}
