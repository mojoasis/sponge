//go:build wireinject
// +build wireinject

package main

import (
	"sponge/internal/api/v1"
	"sponge/internal/dao"
	"sponge/internal/router"
	"sponge/internal/server"
	"sponge/internal/service"

	"github.com/google/wire"
)

func InitApp() (*server.HttpServer, func(), error) {
	panic(wire.Build(
		// 1. Data & DAO
		dao.ProviderSet,

		// 2. Service 层
		service.ProviderSet,

		// 3. API 层
		v1.ProviderSet,

		// 4. 路由层 (新增)
		router.NewRouterGroup,

		// 5. Server 层
		server.NewHttpServer,
	))
}
