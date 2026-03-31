package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine, handler *Handler, store sessions.Store) {
	// Middleware
	router.Use(sessions.Sessions("pizza-tracker", store))

	// Orders
	router.GET("/", handler.ServeNewOrderForm)
	router.POST("/new-order", handler.HandleNewOrderPost)

	// Customers
	router.GET("/customer/:id", handler.serveCustomer)

	// Auth
	router.GET("/login", handler.HandleLoginGet)
	router.POST("/login", handler.HandleLoginPost)
	router.POST("/logout", handler.HandleLogout)

	var admin = router.Group("/admin")
	admin.Use(handler.AuthMiddleware())
	{
		admin.GET("", handler.ServeAdminDashboard)
		admin.POST("/order/:id/update", handler.HandleOrderPut)
		admin.POST("/order/:id/delete", handler.HandleOrderDelete)
	}

	router.Static("/static", "./templates/static")
}
