//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"spring-go-rpc/app/gen/go/system/service/conf"
	"spring-go-rpc/app/system/service/internal/biz"
	"spring-go-rpc/app/system/service/internal/data"
	"spring-go-rpc/app/system/service/internal/server"
	"spring-go-rpc/app/system/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
