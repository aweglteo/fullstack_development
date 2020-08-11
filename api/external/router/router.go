package router

import (
	"github.com/aweglteo/fullstack_development/api/adapter/controllers"
	"github.com/aweglteo/fullstack_development/api/external/postgres"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	Router := gin.Default()

	conn := postgres.GormDB
	// restaurant Controller
	restaurantConroller := controllers.NewRestaurantController(conn)
	Router.POST("/restaurants/stock", func(c *gin.Context) { restaurantConroller.Stock(c) })
	return Router
}
