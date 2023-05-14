// Работа с файлохранилищем
package file

import (
	"encoding/json"
	"os"
	"yandex-devops/storage"
)

// consumer - потребитель для работы с файлом
type consumer struct {
	//file - файл
	file *os.File
	//decoder - дешифратор
	decoder *json.Decoder
}

// NewConsumer - созднаие нового потребителя
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

// Read - получение метрик из файла
func (c *consumer) Read() ([]storage.Metrics, error) {
	metrics := &[]storage.Metrics{}
	if err := c.decoder.Decode(&metrics); err != nil {
		return nil, err
	}
	return *metrics, nil
}

// Close - конец работы с файлом
func (c *consumer) Close() error {
	return c.file.Close()
}
