package main

import (
	"fmt"
	"log/slog"
	"pizza-shop/internal/models"

	"github.com/gin-gonic/gin"
)

type CustomerData struct {
	Title string
	Order models.Order
	Status []string
}

type OrderFormData struct {
	PizzaTypes []string
	PizzaSizes []string
}

type OrderRequest struct {
	Name string			  `form:"name" binding:"required,min=2,max=100"`
	Phone string		  `form:"phone" binding:"required,min=10,max=20"`
	Address string		  `form:"address" binding:"required,min=5,max=200"`
	Types []string		  `form:"types" binding:"required,min=1,dive,valid_pizza_type"`
	Sizes []string		  `form:"sizes" binding:"required,min=1,dive,valid_pizza_size"`
	Instructions []string `form:"instructions" binding:"max=200"`
}

func (this *Handler) ServeNewOrderForm(ctx *gin.Context) {
	ctx.HTML(200, "order.tmpl", OrderFormData{
		PizzaTypes: models.PizzaTypesStr,
		PizzaSizes: models.PizzaSizesStr,
	})
}

func (this *Handler) HandleNewOrderPost(ctx *gin.Context) {
	var form OrderRequest
	var err = ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(400, gin.H { "error": err.Error() })
		return
	}

	fmt.Printf("Form: %+v; Sizes Length: %d\n", form, len(form.Sizes))

	var orderItems = make([]models.OrderItem, len(form.Sizes))
	for i := range orderItems {
		orderItems[i] = models.OrderItem{ Size: form.Sizes[i], Type: form.Types[i], Instructions: form.Instructions[i] }
	}

	var order = models.Order{
		CustomerName: form.Name,
		Phone: form.Phone,
		Address: form.Address,
		Status: models.OrderStatusStr[0],
		Items: orderItems,
	}

	err = this.orders.CreateOrder(&order)
	if err != nil {
		slog.Error("Failed to create order", "error", err)
		ctx.String(500, "Something went wrong: %s", err)
		return
	}

	slog.Info("Order created", "orderId", order.Id, "customer", order.CustomerName)

	fmt.Printf("Order: %+v: ", order)

	ctx.Redirect(303, "/customer/" + order.Id)
}

func (this *Handler) serveCustomer(ctx *gin.Context) {
	var orderId = ctx.Param("id")
	if orderId == "" {
		ctx.String(400, "Order Id is required")
	}

	var order, err = this.orders.GetOrder(orderId)
	if err != nil {
		ctx.String(404, "Order not found")
	}

	ctx.HTML(200, "customer.tmpl", CustomerData {
		Title: "Pizza order status " + orderId,
		Order: *order,
		Status: models.OrderStatusStr,
	})
}
