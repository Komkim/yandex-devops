package config

const (
	pollIntervalDefault   = 2
	reportIntervalDefault = 10

	schemeDefault = "http"
	hostDefault   = "127.0.0.1"
	portDefault   = "8080"
)

type HTTP struct {
	Scheme string
	Port   string
	Host   string
}

type Agent struct {
	Poll   int32
	Report int64
}

type Config struct {
	HTTP
	Agent
}

func Init() *Config {
	return &Config{
		HTTP: HTTP{
			Scheme: schemeDefault,
			Host:   hostDefault,
			Port:   portDefault,
		},
		Agent: Agent{
			Poll:   pollIntervalDefault,
			Report: reportIntervalDefault,
		},
	}
}
