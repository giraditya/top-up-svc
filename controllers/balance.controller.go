package controllers

import (
	"net/http"
	"top-up-service/helpers"
	"top-up-service/presentation"
	"top-up-service/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	c.ShouldBindJSON(&request)

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		helpers.ErrorValidate(c, http.StatusBadRequest, err)
		return
	}

	result, err := u.BalanceService.TopUp(c.Request.Context(), request.UserID, request.Amount)
	if err != nil {
		helpers.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.UserID = result.UserID
	response.Amount = result.Balance

	helpers.SuccessResponse(c, http.StatusOK, "TopUp Success", response)
}
