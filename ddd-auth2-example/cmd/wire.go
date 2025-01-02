//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"learning_tools/ddd-auth2-example/adpter"
	"learning_tools/ddd-auth2-example/domain/aggregate"
	"learning_tools/ddd-auth2-example/domain/service"
	"learning_tools/ddd-auth2-example/infrastructure/conf"
	"learning_tools/ddd-auth2-example/infrastructure/pkg/database/mongo"
	"learning_tools/ddd-auth2-example/infrastructure/pkg/database/redis"
	"learning_tools/ddd-auth2-example/infrastructure/pkg/log"
	"learning_tools/ddd-auth2-example/infrastructure/repository"
)

//go:generate wire
var providerSet = wire.NewSet(
	conf.NewViper,
	conf.NewAppConfigCfg,
	conf.NewLoggerCfg,
	conf.NewRedisConfig,
	conf.NewMongoConfig,
	log.NewLogger,
	redis.NewRedis,
	mongo.NewMongo,
	repository.NewRepository,
	aggregate.NewFactory,
	service.NewService,
	adpter.NewSrv,
)

func NewApp() (*adpter.Server, error) {
	panic(wire.Build(providerSet))
}
