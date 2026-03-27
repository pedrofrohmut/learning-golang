package main

import (
	"log/slog"
	"os"
	"pizza-shop/internal/models"

	"github.com/gin-gonic/gin"
)

func setupLogger() {
	var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func initDb(cfg Config) *models.DbModel {
	var dbModel, err = models.InitDb(cfg.DbPath)
	if err != nil {
		slog.Error("Failed to initialize the database connection", "error", err)
		os.Exit(1)
	}
	slog.Info("Database initialized successfully")
	return dbModel
}

func setupRouterAndRoutes(dbModel *models.DbModel, cfg Config) *gin.Engine {
	var handler = NewHandler(dbModel)
	var router = gin.Default()
	RegisterCustomValidators()
	var err = loadTemplates(router)
	if err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}
	var store = setupSessionStore(dbModel.DB, []byte(cfg.SessionSecretKey))
	setupRoutes(router, handler, store)
	return router
}

func startServer(cfg Config, router *gin.Engine) {
	slog.Info("Server started", "url", "http://localhost:" + cfg.Port)
	router.Run(":" + cfg.Port)
}

func main() {
	setupLogger()
	var cfg = loadConfig()
	var dbModel = initDb(cfg)
	var router = setupRouterAndRoutes(dbModel, cfg)
	startServer(cfg, router)
}
