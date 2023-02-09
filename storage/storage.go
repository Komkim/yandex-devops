package storage

type Storage interface {
	GetOne(key string) (*Metrics, error)
	GetAll() (*[]Metrics, error)
	SetOne(metric Metrics) (*Metrics, error)
	SetAll(metric []Metrics) (*[]Metrics, error)
	Close() error
}
