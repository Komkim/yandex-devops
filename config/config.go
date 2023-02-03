package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/spf13/pflag"
	"time"
)

const (
	defaultHTTPScheme   = "http"
	defaultHTTPPort     = ":8080"
	defaultHTTPHost     = "127.0.0.1"
	defaultHTTPAddress  = "127.0.0.1:8080"
	defaultAgentPoll    = 2 * time.Second
	defaultAgentReport  = 10 * time.Second
	defaultFileInterval = 300 * time.Second
	defaultFilePath     = "/tmp/devops-metrics-db.json"
	defaultFileRestore  = true
)

type HTTP struct {
	Scheme  string `env:"SCHEME" mapstructure:"scheme"`
	Port    string `env:"ADDRESS" mapstructure:"port"`
	Host    string `env:"HOST" mapstructure:"host"`
	Address string `env:"ADDRESS" mapstructure:"address"`
}

type Agent struct {
	Poll   time.Duration `env:"POLL_INTERVAL" mapstructure:"poll"`
	Report time.Duration `env:"REPORT_INTERVAL" mapstructure:"report"`
}

type File struct {
	Interval time.Duration `env:"STORE_INTERVAL" mapstructure:"interval"`
	Path     string        `env:"STORE_FILE" mapstructure:"path"`
	Restore  bool          `env:"RESTORE" mapstructure:"restore"`
}

type Config struct {
	HTTP  HTTP
	Agent Agent
	File  File
}

func InitFlagServer() (*Config, error) {
	cfg := new(Config)

	defaultFlag(cfg)

	initFlagServer(cfg)

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

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

func defaultFlag(cfg *Config) {
	cfg.HTTP.Address = defaultHTTPAddress
	cfg.HTTP.Port = defaultHTTPPort
	cfg.HTTP.Host = defaultHTTPHost
	cfg.HTTP.Scheme = defaultHTTPScheme

	cfg.Agent.Poll = defaultAgentPoll
	cfg.Agent.Report = defaultAgentReport

	cfg.File.Interval = defaultFileInterval
	cfg.File.Path = defaultFilePath
	cfg.File.Restore = defaultFileRestore
}

func initFlagServer(cfg *Config) {
	pflag.StringVarP(&cfg.HTTP.Address, "address", "a", "127.0.0.1:8080", "address")
	pflag.StringVarP(&cfg.File.Path, "file.path", "f", "/tmp/devops-metrics-db.json", "server file path")
	pflag.BoolVarP(&cfg.File.Restore, "file.restore", "r", true, "server file restore")
	pflag.DurationVarP(&cfg.File.Interval, "file.interval", "i", 300*time.Second, "server file report interval")
	pflag.Parse()

}

func initFlagAgent(cfg *Config) {
	pflag.StringVarP(&cfg.HTTP.Address, "address", "a", "127.0.0.1:8080", "address")
	pflag.DurationVarP(&cfg.Agent.Poll, "agent.poll", "p", 2*time.Second, "agent poll interval")
	pflag.DurationVarP(&cfg.Agent.Report, "agent.report", "r", 10*time.Second, "agent report interval")
	pflag.Parse()
}
