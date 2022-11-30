package presentation

type BalanceTopUpRequest struct {
	UserID uint `json:"userid" binding:"required"`
	Amount int  `json:"amount" binding:"required"`
}
