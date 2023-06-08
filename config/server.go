package config

import (
	"encoding/json"
	"errors"
	"github.com/caarlos0/env/v6"
	"github.com/spf13/pflag"
	"os"
	"time"
)

const (
	//defaultFileInterval - значение по умолчанию. Период записи в файл
	defaultFileInterval = 300 * time.Second
	//defaultFilePath - значение по умолчанию. Путь до файла для сохранения метрик
	defaultFilePath = "/tmp/devops-metrics-db.json"
	//defaultFileRestore - значение по умолчанию. Нужно ли записывать в файл.
	defaultFileRestore = true
	//defaultCryptoKeyServer - значение ключа шифрования для сервера
	defaultCryptoKeyServer = ""
)

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

func (d *Duration) UnmarshalText(b []byte) error {
	dr, err := time.ParseDuration(string(b))
	if err != nil {
		return err
	}
	d.Duration = dr
	return nil
}

// Server параметры сервера
type Server struct {
	//FileInterval - интервал записи метрик в файл
	FileInterval Duration `env:"DSN_FILEINTERVAL" mapstructure:"interval" json:"store_interval"`
	//FilePath - путь до файла для записи метрик
	FilePath string `env:"STORE_FILE" mapstructure:"path" json:"store_file"`
	//FileRestore - нужно ли сохранять метрики в файл
	FileRestore bool `env:"RESTORE" mapstructure:"restore" json:"restore"`
	//Key - ключ для работы с хешем
	Key string `env:"KEY" mapstructure:"key" json:"key,omitempty"`
	//DatabaseDSN - параметры работы с базой
	DatabaseDSN string `env:"DATABASE_DSN" mapstructure:"databasedsn" json:"database_dsn"`
	//Address - адрес подключения агента к серверу
	Address string `env:"ADDRESS" mapstructure:"address" json:"address"`
	//CryptoKey - путь до файла с приватным ключом
	CryptoKey string `env:"CRYPTO_KEY" json:"crypto_key"`
	//FileConfig - имя файла конфигурации
	FileConfig string `env:"CONFIG" json:"file_config,omitempty"`
}

// InitFlagServer - инициализация параметров сервера
func InitFlagServer() (*Server, error) {
	cfg := new(Server)

	cfg.parseFlag()

	if len(cfg.FileConfig) > 1 {
		cfgFile, err := loadFileServer(cfg.FileConfig)
		if err != nil {
			return nil, err
		}
		cfg = compareServerConfig(cfg, cfgFile)
	}

	cfgDefault := defaultParamServer()
	cfg = compareServerConfig(cfg, cfgDefault)

	e := &Server{}
	err := env.Parse(e)
	if err != nil {
		return nil, err
	}

	cfg = compareServerConfig(e, cfg)

	return cfg, nil
}

// defaultServer - Присваивание параметров по умолчанию для сервера
func defaultParamServer() *Server {
	s := &Server{}
	s.FileInterval.Duration = defaultFileInterval
	s.FilePath = defaultFilePath
	s.FileRestore = defaultFileRestore
	s.CryptoKey = defaultCryptoKeyServer

	return s
}

// loadFileServer - загружает значения для конфига из файла
func loadFileServer(path string) (s *Server, err error) {
	s = &Server{}
	configFile, err := os.Open(path)
	defer func() {
		err = configFile.Close()
	}()
	if err != nil {
		return nil, err
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&s)
	return s, nil
}

// compareServerConfig - сравнивает два конфига сервера, и записывает в результат все значения из первого и оставшиеся из второго
func compareServerConfig(first, second *Server) *Server {
	result := &Server{}
	result.FileRestore = first.FileRestore

	if len(first.Address) > 0 {
		result.Address = first.Address
	} else if len(second.Address) > 0 {
		result.Address = second.Address
	}

	if len(first.FileConfig) > 0 {
		result.FileConfig = first.FileConfig
	} else if len(second.FileConfig) > 0 {
		result.FileConfig = second.FileConfig
	}

	if len(first.CryptoKey) > 0 {
		result.CryptoKey = first.CryptoKey
	} else if len(second.CryptoKey) > 0 {
		result.CryptoKey = second.CryptoKey
	}

	if len(first.DatabaseDSN) > 0 {
		result.DatabaseDSN = first.DatabaseDSN
	} else if len(second.DatabaseDSN) > 0 {
		result.DatabaseDSN = second.DatabaseDSN
	}

	if len(first.Key) > 0 {
		result.Key = first.Key
	} else if len(second.Key) > 0 {
		result.Key = second.Key
	}

	if len(first.FilePath) > 0 {
		result.FilePath = first.FilePath
	} else if len(second.FilePath) > 0 {
		result.FilePath = second.FilePath
	}

	if first.FileInterval.Duration > 0 {
		result.FileInterval = first.FileInterval
	} else if second.FileInterval.Duration > 0 {
		result.FileInterval = second.FileInterval
	}

	return result
}

// initFlagServer - инициализация переданных параметров для сервера
func (s *Server) parseFlag() {
	pflag.StringVarP(&s.Address, "address", "a", "127.0.0.1:8080", "address")
	pflag.StringVarP(&s.FilePath, "file-path", "f", "/tmp/devops-metrics-db.json", "server file path")
	pflag.BoolVarP(&s.FileRestore, "file-restore", "r", true, "server file restore")
	pflag.DurationVarP(&s.FileInterval.Duration, "file-interval", "i", 300*time.Second, "server file report interval")
	pflag.StringVarP(&s.Key, "key", "k", "", "hash key")
	pflag.StringVarP(&s.DatabaseDSN, "databasedsn", "d", "", "connect postgresql")
	pflag.StringVar(&s.CryptoKey, "crypto-key", "certificat/private.key", "crypto key")
	pflag.StringVarP(&s.FileConfig, "config", "c", "", "path file config")
	pflag.Parse()
}
