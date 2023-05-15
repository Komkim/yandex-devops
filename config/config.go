// config - Пакет с параметрами приложения
package config

import (
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/spf13/pflag"
)

const (
	//defaultHTTPScheme - значение по умолчанию схемы
	defaultHTTPScheme = "http"
	//defaultHTTPAddress - значение по умолчанию адреса запуска сервера
	defaultHTTPAddress = "127.0.0.1:8080"
	//defaultAgentPoll - значение по умолчанию сбора метрик
	defaultAgentPoll = 2 * time.Second
	//defaultAgentReport - значение по умолчанию отправки метрик на сервер
	defaultAgentReport = 10 * time.Second
	//defaultAgentRateLimit - значение по умолчанию запуск воркеров
	defaultAgentRateLimit = 1
	//defaultFileInterval - значение по умолчанию. Период записи в файл
	defaultFileInterval = 300 * time.Second
	//defaultFilePath - значение по умолчанию. Путь до файла для сохранения метрик
	defaultFilePath = "/tmp/devops-metrics-db.json"
	//defaultFileRestore - значение по умолчанию. Нужно ли записывать в файл.
	defaultFileRestore = true
)

var localServerCfg *Config

// HTTP - хранит параметры соединения агента и сервера
type HTTP struct {
	//Scheme - Схема соединения агента к серверу
	Scheme string `env:"SCHEME" mapstructure:"scheme"`
	//Port    string `env:"ADDRESS" mapstructure:"port"`
	//Host    string `env:"HOST" mapstructure:"host"`
	//Address - адрес подключения агента к серверу
	Address string `env:"ADDRESS" mapstructure:"address"`
}

// Agent параметры агента
type Agent struct {
	//Poll - период сбора метрик
	Poll time.Duration `env:"POLL_INTERVAL" mapstructure:"poll"`
	//Report - период отправки метрик на сервер
	Report time.Duration `env:"REPORT_INTERVAL" mapstructure:"report"`
	//Key - ключ для работы с хешем
	Key string `env:"KEY" mapstructure:"key"`
	//RateLimit - колличество воркеров
	RateLimit int `env:"RATE_LIMIT" mapstructure:"rate"`
}

// Server параметры сервера
type Server struct {
	//FileInterval - интервал записи метрик в файл
	FileInterval time.Duration `env:"STORE_INTERVAL" mapstructure:"interval"`
	//FilePath - путь до файла для записи метрик
	FilePath string `env:"STORE_FILE" mapstructure:"path"`
	//FileRestore - нужно ли сохранять метрики в файл
	FileRestore bool `env:"RESTORE" mapstructure:"restore"`
	//Key - ключ для работы с хешем
	Key string `env:"KEY" mapstructure:"key"`
	//DatabaseDSN - параметры работы с базой
	DatabaseDSN string `env:"DATABASE_DSN" mapstructure:"databasedsn"`
}

// Config - хранит параметры сервера, агента и их соединения
type Config struct {
	//HTTP - хранит параметры соединения сервера и агента
	HTTP HTTP
	//Agent - хранит параметры агента
	Agent Agent
	//Server - хранит параметры сервера
	Server Server
}

// InitFlagServer - инициализация параметров сервера
func InitFlagServer() (*Config, error) {
	cfg := new(Config)

	defaultFlag(cfg)

	initFlagServer(cfg)

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	localServerCfg = cfg

	return cfg, nil
}

// InitFlagAgent - инициализация параметров агента
func InitFlagAgent() (*Config, error) {
	cfg := new(Config)

	defaultFlag(cfg)

	initFlagAgent(cfg)

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// defaultFlag - присвоение параметров по умолчанию
func defaultFlag(cfg *Config) {
	cfg.HTTP.Address = defaultHTTPAddress
	//cfg.HTTP.Port = defaultHTTPPort
	//cfg.HTTP.Host = defaultHTTPHost
	cfg.HTTP.Scheme = defaultHTTPScheme

	cfg.Agent.Poll = defaultAgentPoll
	cfg.Agent.Report = defaultAgentReport
	cfg.Agent.RateLimit = defaultAgentRateLimit

	cfg.Server.FileInterval = defaultFileInterval
	cfg.Server.FilePath = defaultFilePath
	cfg.Server.FileRestore = defaultFileRestore
}

// initFlagServer - нициализация переданных параметров для сервера
func initFlagServer(cfg *Config) {
	pflag.StringVarP(&cfg.HTTP.Address, "address", "a", "127.0.0.1:8080", "address")
	pflag.StringVarP(&cfg.Server.FilePath, "file.path", "f", "/tmp/devops-metrics-db.json", "server file path")
	pflag.BoolVarP(&cfg.Server.FileRestore, "file.restore", "r", true, "server file restore")
	pflag.DurationVarP(&cfg.Server.FileInterval, "file.interval", "i", 300*time.Second, "server file report interval")
	pflag.StringVarP(&cfg.Server.Key, "server.key", "k", "", "hash key")
	pflag.StringVarP(&cfg.Server.DatabaseDSN, "server.databasedsn", "d", "", "connect postgresql")
	pflag.Parse()

}

// initFlagAgent - инициализация переданных параметров для агента
func initFlagAgent(cfg *Config) {
	pflag.StringVarP(&cfg.HTTP.Address, "address", "a", "127.0.0.1:8080", "address")
	pflag.DurationVarP(&cfg.Agent.Poll, "agent.poll", "p", 2*time.Second, "agent poll interval")
	pflag.DurationVarP(&cfg.Agent.Report, "agent.report", "r", 10*time.Second, "agent report interval")
	pflag.StringVarP(&cfg.Agent.Key, "agent.key", "k", "", "hash key")
	pflag.IntVarP(&cfg.Agent.RateLimit, "agent.rate_pull", "l", 1, "agent rate limit")
	pflag.Parse()
}

func GetLocalServerCfg() Config {
	return *localServerCfg
}
