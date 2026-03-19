package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbModel struct {
	Order OrderModel
}

func InitDb(dataSourceName string) (*DbModel, error) {
	var db, err = gorm.Open(sqlite.Open(dataSourceName), &gorm.Config {})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)
	}

	err = db.AutoMigrate(&Order {}, &OrderItem {})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate database: %v", err)
	}

	var dbModel = &DbModel { Order: OrderModel { DB: db } }

	return  dbModel, nil
}
