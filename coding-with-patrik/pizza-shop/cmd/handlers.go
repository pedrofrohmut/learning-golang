package main

import "pizza-shop/internal/models"

type Handler struct {
	orders *models.OrderModel
	users *models.UserModel
	notificationManager *NotificationManager
}

func NewHandler(dbModel *models.DbModel) *Handler {
	return &Handler {
		orders: &dbModel.Order,
		users: &dbModel.User,
		notificationManager: NewNotificationManager(),
	}
}
