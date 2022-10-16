package configs

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database
	ServerPort int `envconfig:"SERVER_PORT" default:"80"`
}

type Database struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     int    `envconfig:"DATABASE_PORT" required:"true"`
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envconfig:"DATABASE_NAME" required:"true"`
}

func NewParsedConfig() (Config, error) {
	_ = godotenv.Load(".env")
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	return cnf, err
}
