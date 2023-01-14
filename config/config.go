package config

const (
	pollIntervalDefault   = 2
	reportIntervalDefault = 10

	hostDefault = "127.0.0.1"
	portDefault = "8080"
)

type Http struct {
	Port string
	Host string
}

type Agent struct {
	Poll   int32
	Report int
}

type Config struct {
	Http
	Agent
}

func Init() *Config {
	return &Config{
		Http: Http{
			Host: hostDefault,
			Port: portDefault,
		},
		Agent: Agent{
			Poll:   pollIntervalDefault,
			Report: reportIntervalDefault,
		},
	}
}
