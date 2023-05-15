package file

import (
	"encoding/json"
	"io"
	"os"
	"time"
	"yandex-devops/storage"
)

// producer - потребитель для работы с файлом
type producer struct {
	//file - файл
	file *os.File
	//encoder - кодировщик
	encoder *json.Encoder
}

// NewProducer - создание нового потребителя
func NewProducer(filename string, stream time.Duration) (*producer, error) {
	if stream == 0 {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0777)
		if err != nil {
			return nil, err
		}
		return &producer{
			file:    file,
			encoder: json.NewEncoder(file),
		}, nil
	} else {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return nil, err
		}
		return &producer{
			file:    file,
			encoder: json.NewEncoder(file),
		}, nil
	}
}

// Write - запись метрики в файл
func (p *producer) Write(metric []storage.Metrics) error {
	return p.encoder.Encode(metric)
}

// Close - завершение работы с файлом
func (p *producer) Close() error {
	return p.file.Close()
}

// Cleaning - очистка файла
func (p *producer) Cleaning() error {
	_, err := p.file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	err = p.file.Truncate(0)
	if err != nil {
		return err
	}
	return nil
}
