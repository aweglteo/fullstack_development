package router

import (
	"github.com/aweglteo/fullstack_development/api/adapter/controllers"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()

	// restaurant Controller
	restaurantConroller := controllers.NewRestaurantController()
	Router.POST("/restaurants/stock", func(c *gin.Context) { restaurantConroller.Create(c) })
}
