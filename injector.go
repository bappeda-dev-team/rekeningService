//go:build wireinject
// +build wireinject

package main

import (
	"rekeningService/app"

	"rekeningService/controller"
	"rekeningService/repository"
	"rekeningService/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var rekeningSet = wire.NewSet(
	repository.NewRekeningRepositoryImpl,
	wire.Bind(new(repository.RekeningRepository), new(*repository.RekeningRepositoryImpl)),
	service.NewRekeningServiceImpl,
	wire.Bind(new(service.RekeningService), new(*service.RekeningServiceImpl)),
	controller.NewRekeningControllerImpl,
	wire.Bind(new(controller.RekeningController), new(*controller.RekeningControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		rekeningSet,
		app.NewRouter,
	)
	return nil
}
