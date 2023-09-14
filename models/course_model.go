package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name         string
	Description  string
	Price        float64
	SellingPrice float64
}
