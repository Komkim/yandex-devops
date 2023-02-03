package file

import (
	"sync"
	"yandex-devops/config"
	"yandex-devops/storage"
)

type FileStorage struct {
	mutex    *sync.RWMutex
	producer *producer
	consumer *consumer
}

type FileMetrics struct {
	Metrics []storage.Metrics `json:"metrics_nodes"`
}

func NewFileStorage(cfg *config.Server) (*FileStorage, error) {
	p, err := NewProducer(cfg.FilePath, cfg.FileInterval)
	if err != nil {
		return nil, err
	}
	c, err := NewConsumer(cfg.FilePath)
	if err != nil {
		return nil, err
	}
	return &FileStorage{
		mutex:    &sync.RWMutex{},
		producer: p,
		consumer: c,
	}, nil
}

func (f FileStorage) GetOne(key string) (*storage.Metrics, error) {
	m, err := f.consumer.Read()
	if err != nil {
		return nil, err
	}
	for _, v := range *m {
		if v.ID == key {
			return &v, nil
		}
	}
	return nil, nil
}

func (f FileStorage) GetAll() (*[]storage.Metrics, error) {
	return f.consumer.Read()
}

func (f FileStorage) SetOne(metric storage.Metrics) (*storage.Metrics, error) {
	err := f.producer.Write(&metric)
	if err != nil {
		return nil, err
	}
	return &metric, err

}

func (f FileStorage) SetAll(metric []storage.Metrics) (*[]storage.Metrics, error) {
	metrics := make([]storage.Metrics, 0, len(metric))
	err := f.producer.Cleaning()
	if err != nil {
		return nil, err
	}
	for _, m := range metric {
		mm, err := f.SetOne(m)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, *mm)
	}
	return &metrics, nil
}

func (f *FileStorage) Close() error {
	err := f.producer.Close()
	if err != nil {
		return err
	}
	return f.consumer.Close()
}
