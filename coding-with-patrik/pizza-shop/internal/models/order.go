package models

import (
	"crypto/sha1"
	"time"

	"gorm.io/gorm"
)

// var (
// 	OrderStatus = []string { "Order placed", "Preparing", "Baking", "Quality check", "Ready" }
// )

type OrderStatus string

const (
	OrderPlaced OrderStatus = "Order placed"
	Preparing OrderStatus = "Preparing"
	Baking OrderStatus = "Baking"
	QualityCheck OrderStatus = "Quality check"
	Ready OrderStatus = "Ready"
)

type PizzaType string

const (
	Margheritta PizzaType = "Margheritta"
	Pepperoni PizzaType = "Pepperoni"
	Vegetarian PizzaType = "Vegetarian"
	Hawaiian PizzaType = "Hawaiian"
	BbqChicken PizzaType = "BbqChicken"
	MeatLovers PizzaType = "MeatLovers"
	BuffaloChicken PizzaType = "BuffaloChicken"
	Supreme PizzaType = "Supreme"
	TruffleMushroom PizzaType = "TruffleMushroom"
	FourCheese PizzaType = "FourCheese"
)

type PizzaSize string

const (
	Small PizzaSize = "Small"
	Medium PizzaSize = "Medium"
	Large PizzaSize = "Large"
	XLarge PizzaSize = "XLarge"
)

type OrderModel struct {
	DB *gorm.DB
}

type OrderItem struct {
	Id string `gorm:"primary key; size: 14" json:"id"`
	OrderId string `gorm:"index"`
	Size string `gorm:"not null" json:"size"`
	Type string `gorm:"not null" json:"type"`
	Instructions string `json:"instructions"`
}

type Order struct {
	Id string `gorm:"primary key;size: 14" json:"id"`
	Status string `gorm:"not null" json:"status"`
	CustomerName string `gorm:"not null" json:"customerName"`
	Phone string `gorm:"not null" json:"phone"`
	Address string `gorm:"not null" json:"address"`
	Items []OrderItem `gorm:"foreignKey: OrderId" json:"pizzas"`
	CreatedAt time.Time `json:"createdAt"`
}

func (this *Order) BeforeCreate(tx *gorm.DB) {
	if this.Id == "" {
		this.Id = shortid.
	}
}
