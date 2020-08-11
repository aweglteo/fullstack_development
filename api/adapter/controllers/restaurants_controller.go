package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RestaurantParams struct {
	LinkURL string `json:"linkURL"`
}

type RestaurantController struct {
}

func NewRestaurantController() *RestaurantController {
	return &RestaurantController{}
}

func (controller *RestaurantController) Stock(c *gin.Context) {
	var rp RestaurantParams
	err := c.ShouldBindWith(&rp, binding.Header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "No linkURL data in HTTP Header"})
	}
	c.JSON(http.StatusOK, gin.H{
		"statue":  "sucess",
		"message": "your favorite restaurants is stocked in RDBMS",
	})
}
