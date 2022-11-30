package service

import (
	"context"
	"top-up-service/models"
	"top-up-service/repository"

	"gorm.io/gorm"
)

type BalanceService interface {
	TopUp(ctx context.Context, userid uint, amount int) (models.Balance, error)
}

type balanceService struct {
	DB                       *gorm.DB
	BalanceRepository        repository.BalanceRepository
	BalanceHistoryRepository repository.BalanceHistoryRepository
}

func NewBalanceService(db *gorm.DB, balanceRepository repository.BalanceRepository, balanceHistoryRepository repository.BalanceHistoryRepository) BalanceService {
	return &balanceService{
		DB:                       db,
		BalanceRepository:        balanceRepository,
		BalanceHistoryRepository: balanceHistoryRepository,
	}
}

func (s *balanceService) TopUp(ctx context.Context, userid uint, amount int) (models.Balance, error) {
	db := s.DB.Begin()
	res, err := s.BalanceRepository.FetchByUserID(ctx, db, userid)
	if err != nil {
		return models.Balance{}, err
	}
	newBalance := amount + res.Balance
	err = s.BalanceRepository.Update(ctx, db, newBalance, userid)
	if err != nil {
		db.Rollback()
		return models.Balance{}, err
	}
	err = s.BalanceHistoryRepository.Create(ctx, db, amount, userid)
	if err != nil {
		db.Rollback()
		return models.Balance{}, err
	}
	db.Commit()
	return models.Balance{
		UserID:  userid,
		Balance: newBalance,
	}, nil
}
