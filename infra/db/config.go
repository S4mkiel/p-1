package db

import "fmt"

type Config struct {
	Username string `env:"POSTGRES_USERNAME" required:"true"`
	Password string `env:"POSTGRES_PASSWORD" required:"true"`
	Host     string `env:"POSTGRES_HOST" required:"true"`
	Port     string `env:"POSTGRES_PORT" required:"true"`
	Database string `env:"POSTGRES_DATABASE" required:"true"`
}

func (c Config) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.Host, c.Username, c.Password, c.Database, c.Port,
	)
}
