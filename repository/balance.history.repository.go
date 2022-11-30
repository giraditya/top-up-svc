package repository

import (
	"context"
	"top-up-service/models"

	"gorm.io/gorm"
)

type BalanceHistoryRepository interface {
	Create(ctx context.Context, db *gorm.DB, amount int, userid uint) error
}

type balanceHistoryRepository struct{}

func NewBalanceHistoryRepository() BalanceHistoryRepository {
	return &balanceHistoryRepository{}
}

func (r *balanceHistoryRepository) Create(ctx context.Context, db *gorm.DB, amount int, userid uint) error {
	err := db.Create(&models.BalanceHistory{
		UserID: userid,
		Amount: amount,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
