package db

import (
	"github.com/S4mkiel/p-1/domain/repository"
	"github.com/S4mkiel/p-1/infra/db/source"
	"go.uber.org/fx"
)


var Module = fx.Module(
	"database",
	PostgresModule,
	source.Module,
	fx.Provide(func(src *source.Source) repository.UserRepository { return src.UserSQL }),
	fx.Provide(func(src *source.Source) repository.SexoRepository { return src.SexoSQL }),
)