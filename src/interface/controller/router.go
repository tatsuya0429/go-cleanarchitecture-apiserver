package controller

import (
	"github.com/labstack/echo"
)

// InitRouting Initialize Router
func InitRouting(e *echo.Echo, userController UserController) {
	e.GET("/users/:id", userController.Get())
	e.GET("/users", userController.GetAll())
	e.POST("/users", userController.Post())
	e.PUT("/users/:id", userController.Put())
	e.DELETE("/users/:id", userController.Delete())
}
