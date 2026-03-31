package models

import (
	"time"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderPlaced OrderStatus = "Order placed"
	Preparing OrderStatus = "Preparing"
	Baking OrderStatus = "Baking"
	QualityCheck OrderStatus = "Quality check"
	Ready OrderStatus = "Ready"
)

var OrderStatusStr = []string {
	string(OrderPlaced),
	string(Preparing),
	string(Baking),
	string(QualityCheck),
	string(Ready),
}

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

var PizzaTypesStr = []string {
	string(Margheritta),
	string(Pepperoni),
	string(Vegetarian),
	string(Hawaiian),
	string(BbqChicken),
	string(MeatLovers),
	string(BuffaloChicken),
	string(Supreme),
	string(TruffleMushroom),
	string(FourCheese),
}

type PizzaSize string

const (
	Small PizzaSize = "Small"
	Medium PizzaSize = "Medium"
	Large PizzaSize = "Large"
	XLarge PizzaSize = "XLarge"
)

var PizzaSizesStr = []string {
	string(Small),
	string(Medium),
	string(Large),
	string(XLarge),
}

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

func (this *Order) BeforeCreate(tx *gorm.DB) error {
	if this.Id == "" {
		this.Id = shortid.MustGenerate()
	}

	return nil
}

func (this *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if this.Id == "" {
		this.Id = shortid.MustGenerate()
	}

	return nil
}

func (this *OrderModel) CreateOrder(order *Order) error {
	return this.DB.Create(order).Error
}

func (this *OrderModel) GetOrder(id string) (*Order, error) {
	var order Order
	var err = this.DB.Preload("Items").First(&order, "id = ?", id).Error
	return &order, err
}

func (this *OrderModel) GetAllOrders() ([]Order, error) {
	var orders []Order
	var err = this.DB.Preload("Items").Order("created_at desc").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (this *OrderModel) UpdateOrderStatus(id string, status string) error {
	return this.DB.Model(&Order{}).Where("id = ?", id).Update("status", status).Error;
}

func (this *OrderModel) DeleteOrder(id string) error {
	return this.DB.Select("Items").Where("id = ?", id).Delete(&Order{}).Error
}
