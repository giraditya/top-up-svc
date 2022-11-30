package controllers

import (
	"net/http"
	"top-up-service/presentation"
	"top-up-service/service"

	"github.com/gin-gonic/gin"
)

type BalanceController interface {
	TopUp(c *gin.Context)
}

type balanceController struct {
	BalanceService service.BalanceService
}

func NewBalanceController(balanceService service.BalanceService) BalanceController {
	return &balanceController{
		BalanceService: balanceService,
	}
}

func (u *balanceController) TopUp(c *gin.Context) {
	var request presentation.BalanceTopUpRequest
	var response presentation.BalanceTopUpResponse

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := u.BalanceService.TopUp(c.Request.Context(), request.UserID, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.UserID = result.UserID
	response.Amount = result.Balance

	c.JSON(http.StatusOK, gin.H{"data": response})
}
