//go:build wireinject
// +build wireinject

package main

import (
	"sponge/internal/api/v1"
	"sponge/internal/dao"
	"sponge/internal/server"
	"sponge/internal/service"

	"github.com/google/wire"
)

// InitApp 依赖注入入口
// Wire 会分析依赖关系：Server 依赖 API -> API 依赖 Service -> Service 依赖 DAO
func InitApp() (*server.HttpServer, func(), error) {
	panic(wire.Build(
		dao.NewUserDao,
		service.NewUserService,
		v1.NewUserApi,
		server.NewHttpServer,
	))
}
