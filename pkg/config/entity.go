package config

import (
	"fmt"
	"time"

	"github.com/neoxelox/kit"
	"github.com/neoxelox/kit/util"
	"golang.org/x/text/language"
)

type ConfigService struct {
	Environment         kit.Environment
	Name                string
	Release             string
	TimeZone            time.Location
	GracefulTimeout     time.Duration
	DefaultLocale       language.Tag
	MigrationsPath      string
	TemplatesPath       string
	TemplateFilePattern string
	LocalesPath         string
	LocaleFilePattern   string
	AssetsPath          string
	FilesPath           string
}

type ConfigDatabase struct {
	Host                  string
	Port                  int
	SSLMode               string
	User                  string
	Password              string
	Name                  string
	SchemaVersion         int
	MinConns              int
	MaxConns              int
	MaxConnIdleTime       time.Duration
	MaxConnLifeTime       time.Duration
	DialTimeout           time.Duration
	StatementTimeout      time.Duration
	DefaultIsolationLevel kit.IsolationLevel
}

type ConfigCache struct {
	Host            string
	Port            int
	SSLMode         bool
	Password        string
	MinConns        int
	MaxConns        int
	MaxConnIdleTime time.Duration
	MaxConnLifeTime time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	DialTimeout     time.Duration
}

type ConfigServer struct {
	Host                     string
	Port                     int
	BaseURL                  string
	Origins                  []string
	RequestHeaderMaxSize     int
	RequestBodyMaxSize       int
	RequestFileMaxSize       int
	RequestFilePattern       string
	RequestKeepAliveTimeout  time.Duration
	RequestReadTimeout       time.Duration
	RequestReadHeaderTimeout time.Duration
	ResponseWriteTimeout     time.Duration
}

type ConfigWorker struct {
	Queues         map[string]int
	Concurrency    int
	StrictPriority bool
	StopTimeout    time.Duration
	HealthPort     int
}

type ConfigRunner struct {
}

type ConfigSentry struct {
	DSN string
}

type ConfigGilk struct {
	Port int
}

type ConfigExampleService struct {
	BaseURL string
}

type Config struct {
	Service        ConfigService
	Database       ConfigDatabase
	Cache          ConfigCache
	Server         ConfigServer
	Worker         ConfigWorker
	Runner         ConfigRunner
	Sentry         ConfigSentry
	Gilk           ConfigGilk
	ExampleService ConfigExampleService
}

func NewConfig() *Config {
	return &Config{}
}

func (self Config) String() string {
	return fmt.Sprintf("<Config: %s (%s)>", self.Service.Name, self.Service.Environment)
}

func (self Config) Equals(other Config) bool {
	return util.Equals(self, other)
}

func (self Config) Copy() *Config {
	return util.Copy(self)
}
