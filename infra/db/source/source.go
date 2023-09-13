package source

import (
	src "github.com/S4mkiel/p-1/infra/db/source/postgres"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module(
	"source",
	fx.Provide(NewSQLSources),
)

type Source struct {
	UserSQL *src.UserRepository
	SexoSQL *src.SexoRepository
}

func NewSQLSources(db *gorm.DB) *Source {
	var source = Source{
		UserSQL: src.NewUserRepository(db),
		SexoSQL: src.NewSexoRepository(db),
	}

	return &source
}
