package routes

import (
	"exercise-app/controllers"

	"github.com/labstack/echo/v4"
)

type IndexRoutes string

const INDEX IndexRoutes = ""

func (i IndexRoutes) Build( c *echo.Echo ) {
  routes := c.Group(string(i))
  index_controller := controllers.IndexController{}

  routes.GET("/", index_controller.Show)
  routes.GET("/count", index_controller.AddCount)

  protected := routes.Group("/protected")
  Guard(protected)

  protected.GET("/show", index_controller.Example)
}

