package config

import "time"

type HTTP struct {
	Scheme  string `env:"SCHEME" envDefault:"http"`
	Port    string `env:"ADDRESS" envDefault:"8080"`
	Host    string `env:"HOST" envDefault:"127.0.0.1"`
	Address string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
}

type Agent struct {
	Poll   time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	Report time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
}

type File struct {
	Interval time.Duration `env:"STORE_INTERVAL" envDefault:"10s"`
	Path     string        `env:"STORE_FILE" envDefault:"/tmp/devops-metrics-db.json"`
	Restore  bool          `env:"RESTORE" envDefault:"true"`
}

type Config struct {
	HTTP
	Agent
	File
}
