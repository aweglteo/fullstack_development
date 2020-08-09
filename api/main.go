package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/restaurants/stock", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "tabelog restaurants is stocked in RDBMS",
		})
	})
	r.Run()
}
