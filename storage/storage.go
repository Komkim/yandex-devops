package storage

// Storage - интерфейл для работы с хранилищем
type Storage interface {
	//GetOne - получение метрики
	GetOne(key string) (*Metrics, error)
	//GetAll - получение всех метрик
	GetAll() ([]Metrics, error)
	//SetOne - запись одной метрики
	SetOne(metric Metrics) (*Metrics, error)
	//SetAll - запись нескольх метрик
	SetAll(metric []Metrics) ([]Metrics, error)
	//Close - завершнеие работы с хранилищем
	Close() error
}
