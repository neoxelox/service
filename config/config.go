package config

import (
	"os"
	"strconv"
)

type (
	prometheus struct {
		SubsystemName string
	}

	newRelic struct {
		AppName    string
		LicenseKey string
	}

	sentry struct {
		Dsn string
	}

	// Config stores all the configuration variables for the microservice
	Config struct {
		Prometheus prometheus
		NewRelic   newRelic
		Sentry     sentry
	}
)

// Load setups a new Config instance with all the options
func Load() *Config {
	return &Config{
		Prometheus: prometheus{
			SubsystemName: getEnvAsString("PROMETHEUS_SUBSYSTEM_NAME", "echo"),
		},

		NewRelic: newRelic{
			AppName:    getEnvAsString("NEWRELIC_APPNAME", "NOT_PRESENT"),
			LicenseKey: getEnvAsString("NEWRELIC_LICENSEKEY", "NOT_PRESENT"),
		},

		Sentry: sentry{
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
