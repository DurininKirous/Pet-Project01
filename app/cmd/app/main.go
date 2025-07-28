package main

import (
	"project01/app/internal/db/connection"
	"project01/app/internal/db/config"
	"project01/app/internal/db/users"
	"project01/app/internal/server"
	"project01/app/internal/logging"
	"project01/app/internal/service/monitoring"
	"go.uber.org/zap"
	"context"
	"net/http"
	"log"
)

func main () {
	logger.InitLogger()
	logger.Log.Info("Logger initialized successfully")
	metrics.Init()
	cfg := config.Load()

	ctx := context.Background()
	db := connection.ConnectionStart(ctx, cfg.DB, logger.Log)
	repo := usersDBMethods.New(db, logger.Log)

	defer db.Close()	

	defer logger.Log.Sync()

	router := server.NewRouter(repo, logger.Log)

	logger.Log.Info("Starting HTTP server", zap.String("addr", ":8080"))
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Log.Error("Unable to start server", zap.String("env", "dev"), zap.Error(err))
		log.Fatalf("unable to start server: %v", err)
	}
}
