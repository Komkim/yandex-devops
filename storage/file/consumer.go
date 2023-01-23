package file

import (
	"encoding/json"
	"os"
	"yandex-devops/storage"
)

type consumer struct {
	file    *os.File
	decoder *json.Decoder
}

func NewConsumer(fileName string) (*consumer, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return &consumer{
		file:    file,
		decoder: json.NewDecoder(file),
	}, nil
}
func (c *consumer) Read() (*[]storage.Metrics, error) {
	mm := []storage.Metrics{}
	metrics := storage.Metrics{}
	for c.decoder.More() {
		if err := c.decoder.Decode(&metrics); err != nil {
			return nil, err
		} else {
			mm = append(mm, metrics)
		}
	}
	c.decoder.UseNumber()
	return &mm, nil
}
func (c *consumer) Close() error {
	return c.file.Close()
}
