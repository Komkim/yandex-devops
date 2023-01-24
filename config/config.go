package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHttpScheme   = "http"
	defaultHttpPort     = "8080"
	defaultHttpHost     = "127.0.0.1"
	defaultHttpAddress  = "127.0.0.1:8080"
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
	HTTP
	Agent
	File
}

var (
	agentCmd  = &cobra.Command{}
	serverCmd = &cobra.Command{}
)

func IninServer() (*Config, error) {
	initServer()
	return initCfg()
}

func InitAgent() (*Config, error) {
	initAgent()
	return initCfg()
}

func initCfg() (*Config, error) {

	populateDefaults()

	var cfg Config

	err := unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	err = parseEnv(&cfg)
	if err != nil {
		return nil, err
	}

	err = execute()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func execute() error {
	err := agentCmd.Execute()
	if err != nil {
		return err
	}
	err = serverCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}

func initAgent() {
	agentCmd.PersistentFlags().StringP("http.address", "a", "", "agent http address report")
	agentCmd.Flags().DurationP("agent.poll", "p", 0, "agent poll interval")
	agentCmd.Flags().DurationP("agent.report", "r", 0, "agent report interval")
}

func initServer() {
	agentCmd.PersistentFlags().StringP("http.address", "a", "", "agent http address report")
	agentCmd.Flags().BoolP("file.restore", "r", true, "server file restore")
	agentCmd.Flags().DurationP("file.interval", "i", 0, "server file report interval")
	agentCmd.Flags().StringP("file.path", "f", "", "server file path")
}

func parseEnv(cfg *Config) error {
	err := env.Parse(cfg)
	if err != nil {
		return err
	}
	return nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("agent", &cfg.Agent); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("file", &cfg.File); err != nil {
		return err
	}
	return nil
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.scheme", defaultHttpScheme)
	viper.SetDefault("http.host", defaultHttpHost)
	viper.SetDefault("http.address", defaultHttpAddress)

	viper.SetDefault("agent.poll", defaultAgentPoll)
	viper.SetDefault("agent.report", defaultAgentReport)

	viper.SetDefault("file.path", defaultFilePath)
	viper.SetDefault("file.restore", defaultFileRestore)
	viper.SetDefault("file.interval", defaultFileInterval)
}
