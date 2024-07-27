package server

import "github.com/joeshaw/envdecode"

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
}

func LoadConfig() (Config, error) {
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
