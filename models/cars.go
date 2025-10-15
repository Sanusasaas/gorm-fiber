package models

import "gorm.io/gorm"

type Cars struct {
	ID uint `gorm:"primary key;atoIncrement" json:"id"`
	Brand *string `json:"brand"`
	Name *string `json:"name"`
	Price *float64 `json:"price"`
}

func MigrateCars(db *gorm.DB) error{
	err := db.AutoMigrate(&Cars{})
	return err
}