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
	config.Service.Name = "service-api"
	config.Service.Release = util.GetEnv("SERVICE_API_RELEASE", "wip")
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

	config.Server.Host = util.GetEnv("SERVICE_API_HOST", "localhost")
	config.Server.Port = util.GetEnv("SERVICE_API_PORT", 1111)
	config.Server.BaseURL = util.GetEnv("SERVICE_API_BASE_URL", "http://localhost:1111")
	config.Server.Origins = util.GetEnv("SERVICE_API_ORIGINS", []string{"http://localhost:1111"})
	config.Server.RequestHeaderMaxSize = 1 << 10 // 1 KB
	config.Server.RequestBodyMaxSize = 4 << 10   // 4 KB
	config.Server.RequestFileMaxSize = 2 << 20   // 2 MB
	config.Server.RequestFilePattern = `.*/file.*`
	config.Server.RequestKeepAliveTimeout = config.Service.GracefulTimeout
	config.Server.RequestReadTimeout = config.Service.GracefulTimeout
	config.Server.RequestReadHeaderTimeout = config.Service.GracefulTimeout
	config.Server.ResponseWriteTimeout = config.Service.GracefulTimeout

	config.Sentry.DSN = util.GetEnv("SERVICE_API_SENTRY_DSN", "")

	config.Gilk.Port = util.GetEnv("SERVICE_API_GILK_PORT", 1113)

	config.ExampleService.BaseURL = util.GetEnv("SERVICE_EXAMPLE_SERVICE_BASE_URL", "")

	return config
}
