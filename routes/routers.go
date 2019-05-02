package routers

import (
	"github.com/labstack/echo"
	"github.com/pitsanujiw/restfulgoBasic/controllers"
)

func Router() *echo.Echo {
	// Echo instance
	e := echo.New()
	// Routes
	e.GET("/", controllers.Hello)
	e.GET("/users/:id", controllers.GetUser)
	e.GET("/hello", controllers.GetObject)
	return e
}
