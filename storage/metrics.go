package storage

// Metrics - метрики
type Metrics struct {
	//ID - имя метрики
	ID string `json:"id"` // имя метрики
	//MType - тип
	MType string `json:"type"` // параметр, принимающий значение gauge или counter
	//Delta - значение если тип счетчик
	Delta *int64 `json:"delta,omitempty"` // значение метрики в случае передачи counter
	//Value - значение если это число с плавающей точкой
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
	//Hash - хэш
	Hash string `json:"hash,omitempty"` // значение хеш-функции
}
