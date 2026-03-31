package main

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func (this *Handler) streamSSE(ctx *gin.Context, client chan string) {
	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	ctx.Stream(func (writer io.Writer) bool {
		var msg, ok = <- client
		if ok {
			ctx.SSEvent("message", msg)
			return true
		}
		return false
	})
}

func (this *Handler) notificationHandler(ctx *gin.Context) {
	var orderId = ctx.Query("orderId")

	if orderId == "" {
		fmt.Errorf("notificationHandler: Invalid orderId: %s", orderId)
		ctx.String(400, "Invalid orderId")
		return
	}

	var _, err = this.orders.GetOrder(orderId)
	if err != nil {
		fmt.Errorf("notificationHandler: Order not found with id of %s", orderId)
		ctx.String(404, "Order not found")
		return
	}

	var key = "order:" + orderId
	var client = make(chan string, 10)

	this.notificationManager.AddClient(key, client)

	defer func() {
		this.notificationManager.RemoveClient(key, client)
		slog.Info("Customer client disconnected", "orderId", orderId)
	}()

	this.streamSSE(ctx, client)
}

func (this *Handler) adminNotificationHandler(ctx *gin.Context) {
	var key = "admin:new_orders"
	var client = make(chan string, 10)
	this.notificationManager.AddClient(key, client)

	defer func() {
		this.notificationManager.RemoveClient(key, client)
		slog.Info("Admin client disconnected")
	}()

	this.streamSSE(ctx, client)
}
