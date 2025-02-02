package fxloader

import (
	"ladipage_server/apis/controllers"
	"ladipage_server/apis/middlewares"
	"ladipage_server/apis/resources"
	"ladipage_server/apis/routers"
	"ladipage_server/common/logger"
	"ladipage_server/core/adapters"
	"ladipage_server/core/adapters/repository"
	"ladipage_server/core/services"

	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadEngine()...),
		fx.Options(loadLogger()...),
	}
}

func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(
			adapters.NewPgsql,
		),
		fx.Provide(
			adapters.NewRedis,
		),
		fx.Invoke(func(db *adapters.Pgsql) error {
			return db.Connect()
		}),
		fx.Invoke(func(db *adapters.Redis) error {
			return db.Connect()
		}),
		fx.Provide(repository.NewRepositoryUser),
		fx.Provide(repository.NewRepositoryCache),
		fx.Provide(repository.NewRepositoryTransaction),
		fx.Provide(repository.NewVehicleCategoryRepository),
		fx.Provide(repository.NewRepositoryFileDesc),
		fx.Provide(repository.NewVehicleRepository),
	}
}

func loadUseCase() []fx.Option {
	return []fx.Option{
		fx.Provide(services.NewJwtService),
		fx.Provide(services.NewUserService),
		fx.Provide(services.NewVehicleCategoriesService),
		fx.Provide(services.NewFileDescriptorsService),
		fx.Provide(services.NewVehicleService),
	}
}

func loadEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(routers.NewApiRouter),
		fx.Provide(middlewares.NewMiddlewareCors),
		fx.Provide(controllers.NewUserController),
		fx.Provide(controllers.NewBaseController),
		fx.Provide(resources.NewResource),
		fx.Provide(middlewares.NewMiddlewareJwt),
		fx.Provide(controllers.NewVehicleCategoriesController),
		fx.Provide(controllers.NewFileDescController),
		fx.Provide(controllers.NewVehicleController),
	}
}

func loadLogger() []fx.Option {
	return []fx.Option{
		fx.Provide(logger.NewLogger),
	}
}
