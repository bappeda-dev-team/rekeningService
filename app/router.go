package app

import (
	"rekeningService/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(rekeningController controller.RekeningController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/swagger/doc.json", echoSwagger.WrapHandler)

	e.POST("/rekening", rekeningController.Create)
	e.PUT("/rekening/:id", rekeningController.Update)
	e.DELETE("/rekening/:id", rekeningController.Delete)
	e.GET("/rekening/:id", rekeningController.FindById)
	e.GET("/rekening", rekeningController.FindAll)

	return e
}
