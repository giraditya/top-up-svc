package models

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	UserID  uint
	Balance int `json:"balance"`
}
