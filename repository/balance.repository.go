package repository

import (
	"context"
	"top-up-service/models"

	"gorm.io/gorm"
)

type BalanceRepository interface {
	Update(ctx context.Context, db *gorm.DB, amount int, userid uint) error
	FetchByUserID(ctx context.Context, db *gorm.DB, userid uint) (models.Balance, error)
}

type balanceRepository struct{}

func NewBalanceRepository() BalanceRepository {
	return &balanceRepository{}
}

func (r *balanceRepository) Update(ctx context.Context, db *gorm.DB, amount int, userid uint) error {
	err := db.Where(models.Balance{UserID: userid}).
		Assign(models.Balance{Balance: amount, UserID: userid}).
		FirstOrCreate(&models.Balance{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *balanceRepository) FetchByUserID(ctx context.Context, db *gorm.DB, userid uint) (models.Balance, error) {
	var balance models.Balance
	err := db.Where("user_id =?", userid).Find(&balance).Error
	return balance, err
}
