package server

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/joeshaw/envdecode"
	_ "github.com/lib/pq"
)

type Config struct {
	Port int `env:"APP_PORT"`

	Google struct {
		Host   string `env:"GOOGLE_HOST"`
		APIKey string `env:"GOOGLE_API_KEY"`
	}

	Weather struct {
		Host   string `env:"WEATHER_HOST"`
		APIKey string `env:"WEATHER_API_KEY"`
	}

	OpenAI struct {
		Host   string `env:"OPENAI_HOST"`
		APIKey string `env:"OPENAI_API_KEY"`
	}

	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host                   string `env:"DB_HOST"`
	Username               string `env:"DB_USERNAME"`
	Password               string `env:"DB_PASSWORD"`
	Name                   string `env:"DB_NAME"`
	Port                   string `env:"DB_PORT"`
	IsProd                 bool   `env:"DB_IS_PROD"`
	InstanceConnectionName string `env:"DB_INSTANCE_CONNECTION_NAME"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func ConnectToDB(dbCfg DatabaseConfig) *sqlx.DB {
	var dsn string
	if dbCfg.IsProd {
		dsn = fmt.Sprintf("user=%s password=%s database=%s host=/cloudsql/%s",
			dbCfg.Username, dbCfg.Password, dbCfg.Name, dbCfg.InstanceConnectionName)
	} else {
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			dbCfg.Username, dbCfg.Password, dbCfg.Host,
			dbCfg.Port, dbCfg.Name,
		)
	}

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
