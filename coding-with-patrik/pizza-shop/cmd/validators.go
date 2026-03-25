package main

import (
	"log/slog"
	"pizza-shop/internal/models"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func pizzaTypeValidator(fieldLevel validator.FieldLevel) bool {
	var pizzaType = fieldLevel.Field().String()
	switch pizzaType {
	case string(models.Margheritta),
		 string(models.Pepperoni),
		 string(models.Vegetarian),
		 string(models.Hawaiian),
		 string(models.BbqChicken),
		 string(models.MeatLovers),
		 string(models.BuffaloChicken),
		 string(models.Supreme),
		 string(models.TruffleMushroom),
		 string(models.FourCheese):
		return true
	default:
		return false
	}
}

func pizzaSizeValidator(fieldLevel validator.FieldLevel) bool {
	var pizzaSize = fieldLevel.Field().String()
	switch pizzaSize {
	case string(models.Small),
		 string(models.Medium),
		 string(models.Large),
		 string(models.XLarge):
		return true
	default:
		return false
	}
}

func RegisterCustomValidators() {
	var v, ok = binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("valid_pizza_type", pizzaTypeValidator)
		v.RegisterValidation("valid_pizza_size", pizzaSizeValidator)
		slog.Info("Custom Validation Registered")
	} else {
		slog.Error("Error to register custom validators")
	}
}
