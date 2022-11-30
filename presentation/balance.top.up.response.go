package presentation

type BalanceTopUpResponse struct {
	UserID uint `json:"user_id"`
	Amount int  `json:"amount"`
}
