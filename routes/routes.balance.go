package routes

import (
	"top-up-service/controllers"
	"top-up-service/repository"
	"top-up-service/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitBalanceRoutes(db *gorm.DB, route *gin.Engine) {
	balanceRepository := repository.NewBalanceRepository()
	balanceHistoryRepository := repository.NewBalanceHistoryRepository()
	balanceService := service.NewBalanceService(db, balanceRepository, balanceHistoryRepository)
	balanceController := controllers.NewBalanceController(balanceService)

	groupRoute := route.Group("/v1")
	groupRoute.POST("/balance/topup", balanceController.TopUp)
}
