package main

import (
	"fmt"
	"gin-socmed/config"
	"gin-socmed/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.AuthRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
