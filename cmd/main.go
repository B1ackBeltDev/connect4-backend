package main

import (
	"net/http"
	"os"

	"github.com/B1ackBeltDev/connect4-backend/internals/config"
	"github.com/B1ackBeltDev/connect4-backend/internals/database"
	"github.com/B1ackBeltDev/connect4-backend/internals/logging"
	"github.com/B1ackBeltDev/connect4-backend/internals/routing"
)

func main() {
	logger := logging.NewLogger()

	dbpool, err := database.NewPostgresPool("postgres://sebo:password@localhost:5432/mydatabase")
	if err != nil {
		logger.Error("Failed to initiate Postgres", "error", err.Error())
		os.Exit(1)
	}
	defer dbpool.Close()

	config := config.NewConfig()

	router := routing.NewRouter()

	server := http.Server{
		Addr:    config.Server.IP + ":" + config.Server.Port,
		Handler: router,
	}

	logger.Info("Starting server", "IP", config.Server.IP, "Port", config.Server.Port)

	if err := server.ListenAndServe(); err != nil {
		logger.Error("Failed to start server. Exiting.")
	}
}
