package config

import (
	"encoding/json"
	"github.com/caarlos0/env/v6"
	"github.com/spf13/pflag"
	"os"
	"time"
)

const (
	//defaultHTTPScheme - значение по умолчанию схемы
	defaultHTTPScheme = "https"
	//defaultHTTPAddress - значение по умолчанию адреса запуска сервера
	defaultHTTPAddress = "127.0.0.1:8080"
	//defaultAgentPoll - значение по умолчанию сбора метрик
	defaultAgentPoll = 2 * time.Second
	//defaultAgentReport - значение по умолчанию отправки метрик на сервер
	defaultAgentReport = 10 * time.Second
	//defaultAgentRateLimit - значение по умолчанию запуск воркеров
	defaultAgentRateLimit = 1
	//defaultCryptoKeyAgent - значение ключа шифрования для агента
	defaultCryptoKeyAgent = ""
	//defaultApiType - дефолтное значение типа запускаемого клиента
	defaultApiType = REST

	GRPC = "GRPC"
	REST = "REST"
)

// Agent параметры агента
type Agent struct {
	//Poll - период сбора метрик
	Poll Duration `env:"POLL_INTERVAL" mapstructure:"poll" json:"poll_interval"`
	//Report - период отправки метрик на сервер
	Report Duration `env:"REPORT_INTERVAL" mapstructure:"report" json:"report_interval"`
	//Key - ключ для работы с хешем
	Key string `env:"KEY" mapstructure:"key" json:"key,omitempty"`
	//RateLimit - колличество воркеров
	RateLimit int `env:"RATE_LIMIT" mapstructure:"rate" json:"rate_limit,omitempty"`
	//Scheme - Схема соединения агента к серверу
	Scheme string `env:"SCHEME" mapstructure:"scheme" json:"scheme,omitempty"`
	//Address - адрес подключения агента к серверу
	Address string `env:"ADDRESS" mapstructure:"address" json:"address"`
	//CryptoKey - путь до файла с приватным ключом
	CryptoKey string `env:"CRYPTO_KEY" json:"crypto_key"`
	//FileConfig - имя файла конфигурации
	FileConfig string `env:"CONFIG" json:"file_config,omitempty"`
	//AptType - тип запускаемого агента
	ApiType string `env:"APi_TYPE" json:"api_type"`
}

// InitFlagAgent - инициализация параметров агента
func InitFlagAgent() (*Agent, error) {
	cfg := new(Agent)

	cfg.parseFlag()
	if len(cfg.FileConfig) > 1 {
		cfgFile, err := loadFileAgent(cfg.FileConfig)
		if err != nil {
			return nil, err
		}
		cfg = compareAgentConfig(cfg, cfgFile)
	}

	cfgDefault := defaultParamAgent()
	cfg = compareAgentConfig(cfg, cfgDefault)

	e := &Agent{}
	err := env.Parse(e)
	if err != nil {
		return nil, err
	}

	cfg = compareAgentConfig(e, cfg)
	return cfg, nil
}

// defaultAgent - Присваивание параметров по умолчанию для агента
func defaultParamAgent() *Agent {
	a := &Agent{}
	a.Address = defaultHTTPAddress
	a.Scheme = defaultHTTPScheme
	a.CryptoKey = defaultCryptoKeyAgent
	a.Poll.Duration = defaultAgentPoll
	a.Report.Duration = defaultAgentReport
	a.RateLimit = defaultAgentRateLimit
	a.ApiType = defaultApiType
	return a
}

// loadFileAgent - загружает значения для конфига из файла
func loadFileAgent(path string) (a *Agent, err error) {
	a = &Agent{}
	configFile, err := os.Open(path)
	defer func() {
		err = configFile.Close()
	}()
	if err != nil {
		return nil, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&a)
	return a, nil
}

// compareAgentConfig - сравнивает два конфига агента, и записывает в результат все значения из первого и оставшиеся из второго
func compareAgentConfig(first, second *Agent) *Agent {
	result := &Agent{}

	if len(first.Address) > 0 {
		result.Address = first.Address
	} else if len(second.Address) > 0 {
		result.Address = second.Address
	}
	if first.Poll.Duration > 0 {
		result.Poll = first.Poll
	} else if second.Poll.Duration > 0 {
		result.Poll = second.Poll
	}
	if first.Report.Duration > 0 {
		result.Report = first.Report
	} else if second.Report.Duration > 0 {
		result.Report = second.Report
	}
	if len(first.Key) > 0 {
		result.Key = first.Key
	} else if len(second.Key) > 0 {
		result.Key = second.Key
	}
	if first.RateLimit > 0 {
		result.RateLimit = first.RateLimit
	} else if second.RateLimit > 0 {
		result.RateLimit = second.RateLimit
	}
	if len(first.CryptoKey) > 0 {
		result.CryptoKey = first.CryptoKey
	} else if len(second.CryptoKey) > 0 {
		result.CryptoKey = second.CryptoKey
	}
	if len(first.FileConfig) > 0 {
		result.FileConfig = first.FileConfig
	} else if len(second.FileConfig) > 0 {
		result.FileConfig = second.FileConfig
	}
	if len(first.Scheme) > 0 {
		result.Scheme = first.Scheme
	} else if len(second.Scheme) > 0 {
		result.Scheme = second.Scheme
	}
	if len(first.ApiType) > 0 {
		result.ApiType = first.ApiType
	} else if len(second.ApiType) > 0 {
		result.ApiType = second.ApiType
	}

	return result
}

// initFlagAgent - инициализация переданных параметров для агента
func (a *Agent) parseFlag() {
	pflag.StringVarP(&a.Address, "address", "a", "127.0.0.1:8080", "address")
	pflag.DurationVarP(&a.Poll.Duration, "poll", "p", 2*time.Second, "agent poll interval")
	pflag.DurationVarP(&a.Report.Duration, "report", "r", 10*time.Second, "agent report interval")
	pflag.StringVarP(&a.Key, "key", "k", "", "hash key")
	pflag.IntVarP(&a.RateLimit, "rate-limit", "l", 1, "agent rate limit")
	pflag.StringVar(&a.CryptoKey, "crypto-key", "", "crypto key")
	pflag.StringVarP(&a.FileConfig, "config", "c", "", "path file config")
	pflag.StringVarP(&a.ApiType, "api-type", "t", "REST", "api type")
	pflag.Parse()
}
