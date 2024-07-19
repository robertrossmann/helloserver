package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"helloserver/environ"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

func init() {
	// An actual deployment. Do not touch any of the .env files - configuration must come from the actual environment ⚠️
	if _, found := os.LookupEnv("APP_ENV"); found {
		return
	}

	pkgroot := environ.PackageRoot()
	// It's ok if this file is missing
	_ = godotenv.Load(filepath.Join(pkgroot, ".env"))
	// but this one is required
	err := godotenv.Load(filepath.Join(pkgroot, ".env.local"))
	if err != nil {
		panic(fmt.Errorf("loading .env.local: %w", err))
	}
}

type Config struct {
	Port                  int        `env:"PORT" envDefault:"3000"`
	LogLevel              slog.Level `env:"LOG_LEVEL" envDefault:"info"`
	ServerShutdownTimeout time.Duration
}

// New loads the configuration from the environment and returns a Config struct ready for use
func New() (Config, error) {
	config := Config{
		ServerShutdownTimeout: 15 * time.Second,
	}

	if err := env.Parse(&config); err != nil {
		return Config{}, fmt.Errorf("marshalling env to config: %w", err)
	}

	return config, nil
}
