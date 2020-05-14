package server

import (
	"fmt"
	"os"
	"strconv"
)

type (
	_app struct {
		Port    int
		Debug   bool
		Name    string
		Version string
		Release string
	}

	_database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
		Dsn      string
	}

	_newRelic struct {
		LicenseKey string
	}

	_sentry struct {
		Dsn string
	}

	// Configuration stores all the settings for the application
	Configuration struct {
		App      *_app
		Database *_database
		NewRelic *_newRelic
		Sentry   *_sentry
	}
)

// NewConfiguration creates a new Configuration instance
func NewConfiguration() (*Configuration, error) {
	return &Configuration{
		App: &_app{
			Port:    getEnvAsInt("APP_PORT", 8000),
			Debug:   getEnvAsBool("APP_DEBUG", false),
			Name:    getEnvAsString("APP_NAME", "mst"),
			Version: getEnvAsString("APP_VERSION", "v1"),
			Release: getEnvAsString("APP_RELEASE", "master"),
		},

		Database: &_database{
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

		NewRelic: &_newRelic{
			LicenseKey: getEnvAsString("NEWRELIC_LICENSEKEY", "NOT_PRESENT"),
		},

		Sentry: &_sentry{
			Dsn: getEnvAsString("SENTRY_DSN", "NOT_PRESENT"),
		},
	}, nil
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
