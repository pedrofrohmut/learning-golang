package main

import (
	"ecommerce/internal/products"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type DbConfig struct {
	dsn string
}

type Config struct {
	addr string
	db DbConfig
}

type Application struct {
	config Config
}

// Mount
func (app *Application) mount() http.Handler {
	var router = chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/health", func(writer http.ResponseWriter, router *http.Request) {
		writer.Write([]byte("Server is on."))
	})

	var productsService = products.NewSvc()
	var productsHandler = products.NewHandler(productsService)
	router.Get("/products", productsHandler.ListProducts)

	return router
}

// Run
func (app *Application) run(handler http.Handler) error {
	var server = &http.Server {
		Addr: app.config.addr,
		Handler: handler,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Second * 60,
	}

	return server.ListenAndServe()
}
