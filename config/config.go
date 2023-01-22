package config

import "time"

const (
	pollIntervalDefault   = 2
	reportIntervalDefault = 10

	schemeDefault = "http"
	hostDefault   = "127.0.0.1"
	portDefault   = "8080"
)

type HTTP struct {
	Scheme string `env:"SCHEME" envDefault:"http"`
	Port   string `env:"ADDRESS" envDefault:"8080"`
	Host   string `env:"HOST" envDefault:"127.0.0.1"`
}

type Agent struct {
	Poll   int32 `env:"POLL_INTERVAL" envDefault:"2"`
	Report int64 `env:"REPORT_INTERVAL" envDefault:"10"`
}

type Config struct {
	Scheme  string        `env:"SCHEME" envDefault:"http"`
	Address string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	Poll    time.Duration `env:"POLL_INTERVAL" envDefault:"2s"`
	Report  time.Duration `env:"REPORT_INTERVAL" envDefault:"10s"`
}

//type Config struct {
//	HTTP
//	Agent
//	Address string `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
//}

//func Init() *Config {
//	return &Config{
//		HTTP: HTTP{
//			Scheme: schemeDefault,
//			Host:   hostDefault,
//			Port:   portDefault,
//		},
//		Agent: Agent{
//			Poll:   pollIntervalDefault,
//			Report: reportIntervalDefault,
//		},
//	}
//}
