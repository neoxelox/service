package config

import (
	"fmt"
	"os"
	"strconv"
)

type (
	app struct {
		Port    int
		Debug   bool
		Name    string
		Release string
	}

	database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		Dsn      string
	}

	newRelic struct {
		LicenseKey string
	}

	sentry struct {
		Dsn string
	}

	// Config stores all the configuration variables for the microservice
	Config struct {
		App      *app
		Database *database
		NewRelic *newRelic
		Sentry   *sentry
	}
)

// Load setups a new Config instance with all the options
func Load() *Config {
	return &Config{
		App: &app{
			Port:    getEnvAsInt("APP_PORT", 8000),
			Debug:   getEnvAsBool("APP_DEBUG", false),
			Name:    getEnvAsString("APP_NAME", "mst"),
			Release: getEnvAsString("APP_RELEASE", "v1"),
		},

		Database: &database{
			Host:     getEnvAsString("DATABASE_HOST", "localhost"),
			Port:     getEnvAsInt("DATABASE_PORT", 5432),
			User:     getEnvAsString("DATABASE_USER", "postgres"),
			Password: getEnvAsString("DATABASE_PASSWORD", "postgres"),
			Dsn: fmt.Sprintf("postgresql://%s:%s@%s:%d",
				getEnvAsString("DATABASE_USER", "postgres"),
				getEnvAsString("DATABASE_PASSWORD", "postgres"),
				getEnvAsString("DATABASE_HOST", "localhost"),
				getEnvAsInt("DATABASE_PORT", 5432),
			),
		},

		NewRelic: &newRelic{
			LicenseKey: getEnvAsString("NEWRELIC_LICENSEKEY", "NOT_PRESENT"),
		},

		Sentry: &sentry{
			Dsn: getEnvAsString("SENTRY_DSN", "NOT_PRESENT"),
		},
	}
}

func getEnvAsString(key string, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return def
}

func getEnvAsInt(key string, def int) int {
	valueStr := getEnvAsString(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return def
}

func getEnvAsBool(key string, def bool) bool {
	valueStr := getEnvAsString(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return def
}
