package presentation

type BalanceTopUpRequest struct {
	UserID uint `json:"userid" validate:"required"`
	Amount int  `json:"amount" validate:"required"`
}
