package models

import "gorm.io/gorm"

type BalanceHistory struct {
	gorm.Model
	UserID uint
	Amount int
}
