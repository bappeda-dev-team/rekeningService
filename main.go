package main

import (
	"fmt"
	"rekeningService/docs"
	"rekeningService/helper"

	"os"

	"github.com/labstack/echo/v4"
)

func NewServer(e *echo.Echo) *echo.Echo {
	return e
}

// @title Rekening Service
// @version 1.0
// @description API Data Master Rekening Kab Bontang
// @termsOfService http://swagger.io/terms/

// @contact.name KK-DevTeam
// @contact.email bappedadevteam@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host ${PROD_HOSTNAME}
// @BasePath /

func main() {

	// DEPRECATED jalankan flyway secara terpisah
	// app.RunFlyway()

	server := InitializedServer()
	host := os.Getenv("host")
	port := os.Getenv("port")
	prod := os.Getenv("PROD_HOSTNAME")

	docs.SwaggerInfo.Host = fmt.Sprintf("%v", prod)

	addr := fmt.Sprintf("%s:%s", host, port)

	err := server.Start(addr)
	helper.PanicIfError(err)
}
