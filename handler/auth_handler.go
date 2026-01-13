package handler

import (
	"gin-socmed/dto"
	"gin-socmed/errorhandler"
	"gin-socmed/helper"
	"gin-socmed/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (*authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register successfully, please login",
	})

	c.JSON(http.StatusCreated, res)
}
