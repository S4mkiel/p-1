package service

import (
	"github.com/S4mkiel/p-1/domain/service"
	"go.uber.org/fx"
)



var Module = fx.Module(
	"services",
	fx.Provide(service.NewUserService),
	fx.Provide(service.NewSexoService),
)