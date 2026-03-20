package main

import (
	"log/slog"
	"os"
	"pizza-shop/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Logger
	var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	slog.Info("Start working")

	var cfg = loadConfig()

	// Database connection
	var dbModel, err = models.InitDb(cfg.DbPath)
	if err != nil {
		slog.Error("Failed to initialize the database connection", "error", err)
		os.Exit(1)
	}
	slog.Info("Database initialized successfully")

	// Setup routes and start the server
	var handler = NewHandler(dbModel)
	var router = gin.Default()
	err = loadTemplates(router)
	if err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}
	setupRoutes(router, handler)
	slog.Info("Server started", "url", "http://localhost:" + cfg.Port)
	router.Run(":" + cfg.Port)
}
