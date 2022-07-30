//go:build wireinject
// +build wireinject

package main

import (
	"service/api/rest/internal/handler"
	"service/api/rest/internal/logic"
	userLogic "service/internal/user/logic"
	"service/internal/user/repo"

	"infra"

	"github.com/google/wire"
)

func initApp() *App {
	wire.Build(
		handler.NewHandler,
		logic.NewCreateUserLogic,
		logic.NewGetUserLogic,
		userLogic.NewLogic,
		repo.NewMySQLUserRepo,
		infra.ProvideMySQL,
		wire.Struct(new(App), "*"),
	)

	return &App{}
}
