package main

import (
	"runtime"
	"time"

	"github.com/neoxelox/kit"
	"github.com/neoxelox/kit/util"
	"golang.org/x/text/language"

	"service/pkg/config"
)

func NewConfig() *config.Config {
	config := config.NewConfig()

	config.Service.Environment = kit.Environment(util.GetEnv("SERVICE_ENVIRONMENT", "dev"))
	config.Service.Name = "service-worker"
	config.Service.Release = util.GetEnv("SERVICE_WORKER_RELEASE", "wip")
	config.Service.TimeZone = *time.UTC
	config.Service.GracefulTimeout = 30 * time.Second
	config.Service.DefaultLocale = language.English
	config.Service.MigrationsPath = "migrations"
	config.Service.TemplatesPath = "templates"
	config.Service.TemplateFilePattern = `^.*\.(html|txt|md)$`
	config.Service.LocalesPath = "locales"
	config.Service.LocaleFilePattern = `^.*\.(yml|yaml)$`
	config.Service.AssetsPath = "assets"
	config.Service.FilesPath = "files"

	config.Database.Host = util.GetEnv("SERVICE_DATABASE_HOST", "postgres")
	config.Database.Port = util.GetEnv("SERVICE_DATABASE_PORT", 5432)
	config.Database.SSLMode = util.GetEnv("SERVICE_DATABASE_SSLMODE", "disable")
	config.Database.User = util.GetEnv("SERVICE_DATABASE_USER", "service")
	config.Database.Password = util.GetEnv("SERVICE_DATABASE_PASSWORD", "service")
	config.Database.Name = util.GetEnv("SERVICE_DATABASE_NAME", "service")
	config.Database.SchemaVersion = 2
	config.Database.MinConns = 1
	config.Database.MaxConns = max(4, 2*runtime.GOMAXPROCS(-1))
	config.Database.MaxConnIdleTime = 30 * time.Minute
	config.Database.MaxConnLifeTime = 1 * time.Hour
	config.Database.DialTimeout = config.Service.GracefulTimeout
	config.Database.StatementTimeout = config.Service.GracefulTimeout
	config.Database.DefaultIsolationLevel = kit.IsoLvlReadCommitted

	config.Cache.Host = util.GetEnv("SERVICE_CACHE_HOST", "redis")
	config.Cache.Port = util.GetEnv("SERVICE_CACHE_PORT", 6379)
	config.Cache.SSLMode = util.GetEnv("SERVICE_CACHE_SSLMODE", false)
	config.Cache.Password = util.GetEnv("SERVICE_CACHE_PASSWORD", "redis")
	config.Cache.MinConns = 1
	config.Cache.MaxConns = max(8, 4*runtime.GOMAXPROCS(-1))
	config.Cache.MaxConnIdleTime = 30 * time.Minute
	config.Cache.MaxConnLifeTime = 1 * time.Hour
	config.Cache.ReadTimeout = config.Service.GracefulTimeout
	config.Cache.WriteTimeout = config.Service.GracefulTimeout
	config.Cache.DialTimeout = config.Service.GracefulTimeout

	config.Worker.Queues = map[string]int{
		"critical":   6,
		"default":    3,
		"irrelevant": 1,
	}
	config.Worker.Concurrency = 4 * runtime.GOMAXPROCS(-1)
	config.Worker.StrictPriority = false
	config.Worker.StopTimeout = config.Service.GracefulTimeout
	config.Worker.HealthPort = util.GetEnv("SERVICE_WORKER_HEALTH_PORT", 1112)

	config.Sentry.DSN = util.GetEnv("SERVICE_WORKER_SENTRY_DSN", "")

	config.ExampleService.BaseURL = util.GetEnv("SERVICE_EXAMPLE_SERVICE_BASE_URL", "")

	return config
}
