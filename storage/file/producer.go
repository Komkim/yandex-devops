package file

import (
	"encoding/json"
	"io"
	"os"
	"yandex-devops/storage"
)

type producer struct {
	file    *os.File
	encoder *json.Encoder
}

func NewProducer(filename string) (*producer, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}

	return &producer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}
func (p *producer) Write(metric *storage.Metrics) error {
	return p.encoder.Encode(&metric)
}
func (p *producer) Close() error {
	return p.file.Close()
}

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
