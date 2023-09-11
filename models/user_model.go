package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string  `jsong:"name" gorm:"text;not null;default: null`
	LastName    string  `json:"last_name" gorm:"text;not null;default: null`
	PhoneNumber string  `json:"phone_number" gorm:"text;not null;default: null`
	Email       string  `json:"phone_number" gorm:"text;not null;default: null`
	Balance     float64 `json:"blance" gorm:"decimal;not null;default: 0`
	Birthdate   string  `json:"birthdate" gorm:"timestamp with time zone;not null;`
	Password    string  `gorm:"text;not null;default: nul`
}
