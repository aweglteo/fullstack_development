package router

import (
	"github.com/aweglteo/fullstack_development/api/adapter/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	Router := gin.Default()

	// restaurant Controller
	restaurantConroller := controllers.NewRestaurantController()
	Router.POST("/restaurants/stock", func(c *gin.Context) { restaurantConroller.Stock(c) })
	return Router
}
