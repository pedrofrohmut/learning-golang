package main

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine, handler *Handler) {
	router.GET("/", handler.ServeNewOrderForm)
	router.POST("/orders", handler.HandleNewOrderPost)
	router.GET("/orders/:id", handler.serveCustomer)

	router.Static("/static", "/templates/static")
}
