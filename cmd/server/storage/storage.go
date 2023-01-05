package storage

type Storage interface {
	GetOne(key string) (Metric, error)
	GetAll() ([]Metric, error)
	SetOne(metric Metric) error
	SetAll(metric []Metric) error
}

type Repositories struct {
	Storage Storage
}

func NewRepositories(s Storage) *Repositories {
	return &Repositories{s}
}
