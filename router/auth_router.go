package router

import (
	"gin-socmed/config"
	"gin-socmed/handler"
	"gin-socmed/repository"
	"gin-socmed/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	autRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(autRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
}
