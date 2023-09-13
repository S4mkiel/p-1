package db

import (
	"context"
	"log"

	"github.com/S4mkiel/p-1/domain/entity"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresModule = fx.Module(
	"postgres",
	fx.Provide(NewClient),
	fx.Invoke(HookDatabase),
	fx.Invoke(migrate),
)

func NewClient(cfg Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.ConnectionString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func HookDatabase(lc fx.Lifecycle, db *gorm.DB, logger *zap.SugaredLogger) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			dbDriver, err := db.DB()
			if err != nil {
				logger.Fatal("Failed to get database driver", zap.Error(err))
				return err
			}

			err = dbDriver.Ping()
			if err != nil {
				logger.Fatal("Failed to ping database", zap.Error(err))
				return err
			}

			logger.Info("Database OK!")
			return nil
		},
		OnStop: func(context.Context) error {
			dbDriver, err := db.DB()
			if err != nil {
				logger.Fatal("Failed to get database driver", zap.Error(err))
				return err
			}

			err = dbDriver.Close()
			if err != nil {
				logger.Fatal("Failed to close database", zap.Error(err))
				return err
			}

			logger.Info("Database connection closed!")
			return nil
		},
	})
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.Sexo{},
		&entity.User{},
	)
	
	if err != nil {
		log.Panicln(err)
	}
}
