package myclient

const certFile = "certificat/certificate.crt"

// Metrics - метрики
type Metrics struct {
	//ID - имя метрики
	ID string `json:"id"` // имя метрики
	//MType - тим метрики
	MType string `json:"type"` // параметр, принимающий значение gauge или counter
	//Delta - значение метрики в случае передачи счетчика
	Delta *int64 `json:"delta,omitempty"` // значение метрики в случае передачи counter
	//Value - значение метрики в случае передачи числа с плавающей точкой
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
	//Hash - значение хэш-функции
	Hash string `json:"hash,omitempty"` // значение хеш-функции
}

type Client interface {
	SendOneMetric(metric Metrics) error
	SendAllMetric(metrics []Metrics) error
	//SendOneMetric(metric Metrics) interface{}
}
