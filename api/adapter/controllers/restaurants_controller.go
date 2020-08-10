package controllers

import "github.com/gin-gonic/gin"

type RestaurantController struct {
}

func NewRestaurantController() *RestaurantController {
	return &RestaurantController{}
}

func (controller *RestaurantController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "your favorite restaurants is stocked in RDBMS",
	})
}
