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
	Key    string        `env:"KEY" mapstructure:"key"`
}

type Server struct {
	FileInterval time.Duration `env:"STORE_INTERVAL" mapstructure:"interval"`
	FilePath     string        `env:"STORE_FILE" mapstructure:"path"`
	FileRestore  bool          `env:"RESTORE" mapstructure:"restore"`
	Key          string        `env:"KEY" mapstructure:"key"`
	DatabaseDSN  string        `env:"DATABASE_DSN" mapstructure:"databasedsn"`
}

type Config struct {
	HTTP   HTTP
	Agent  Agent
	Server Server
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

	cfg.Server.FileInterval = defaultFileInterval
	cfg.Server.FilePath = defaultFilePath
	cfg.Server.FileRestore = defaultFileRestore
}

func initFlagServer(cfg *Config) {
	pflag.StringVarP(&cfg.HTTP.Address, "address", "a", "127.0.0.1:8080", "address")
	pflag.StringVarP(&cfg.Server.FilePath, "file.path", "f", "/tmp/devops-metrics-db.json", "server file path")
	pflag.BoolVarP(&cfg.Server.FileRestore, "file.restore", "r", true, "server file restore")
	pflag.DurationVarP(&cfg.Server.FileInterval, "file.interval", "i", 300*time.Second, "server file report interval")
	pflag.StringVarP(&cfg.Server.Key, "server.key", "k", "", "hash key")
	pflag.StringVarP(&cfg.Server.DatabaseDSN, "server.databasedsn", "d", "", "connect postgresql")
	pflag.Parse()

}

func initFlagAgent(cfg *Config) {
	pflag.StringVarP(&cfg.HTTP.Address, "address", "a", "127.0.0.1:8080", "address")
	pflag.DurationVarP(&cfg.Agent.Poll, "agent.poll", "p", 2*time.Second, "agent poll interval")
	pflag.DurationVarP(&cfg.Agent.Report, "agent.report", "r", 10*time.Second, "agent report interval")
	pflag.StringVarP(&cfg.Agent.Key, "agent.key", "k", "", "hash key")
	pflag.Parse()
}
