package backend

import (
	"fmt"
	"os"
	"path/filepath"

	"helloserver/environ"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port int `env:"PORT" envDefault:"3000"`
}

func LoadDotenv() {
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
		panic(err)
	}
}

// NewConfig loads the configuration from the environment and returns a Config struct ready for use
func NewConfig() (Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		return Config{}, fmt.Errorf("marshalling env to config: %w", err)
	}

	return config, nil
}
