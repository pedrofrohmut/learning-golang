package main

import "pizza-shop/internal/models"

type Handler struct {
	orders *models.OrderModel
}

func NewHandler(dbModel *models.DbModel) *Handler {
	return &Handler {
		orders: &dbModel.Order,
	}
}
