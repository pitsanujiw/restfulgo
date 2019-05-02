package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pitsanujiw/restfulgoBasic/models"
)

// Handler
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler
func GetObject(c echo.Context) error {
	u := &models.Users{
		Name:  "Jon",
		Email: "jon@labstack.com",
	}
	return c.JSON(http.StatusOK, u)
}

//
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	fmt.Println(c)
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
