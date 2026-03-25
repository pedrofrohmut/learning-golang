package main

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine, handler *Handler) {
	// Orders
	router.GET("/", handler.ServeNewOrderForm)
	router.POST("/new-order", handler.HandleNewOrderPost)

	// Customers
	router.GET("/customer/:id", handler.serveCustomer)

	router.Static("/static", "./templates/static")
}
