package controllers

import (
	"net/http"

	"github.com/aweglteo/fullstack_development/api/adapter/gateway"
	"github.com/aweglteo/fullstack_development/api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

type RestaurantParams struct {
	LinkURL string `json:"linkURL"`
}

type RestaurantController struct {
	Interactor usecase.RestaurantInteractor
}

func NewRestaurantController(conn *gorm.DB) *RestaurantController {
	return &RestaurantController{
		Interactor: usecase.RestaurantInteractor{
			RestaurantRepository: &gateway.RestaurantRepository{
				Conn: conn,
			},
		},
	}
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
