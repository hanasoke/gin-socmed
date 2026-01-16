package middleware

import (
	"gin-socmed/errorhandler"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorhandler.HandleError(c, &errorhandler.UnathorizedError{Message: "Unauthorize"})
			c.Abort()
			return
		}

	}
}
