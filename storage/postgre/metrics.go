package postgresql

import (
	"database/sql"
	"time"
)

// metrics - Метрики для базы
type metrics struct {
	//ID - идентификатор
	ID string `db:"id"`
	//Name - название
	Name string `db:"name"`
	//MType - тип
	MType string `db:"type"`
	//Delta - значение если это счетчик
	Delta sql.NullInt64 `db:"delta"`
	//Value - значение если это число с плавающей точкой
	Value sql.NullFloat64 `db:"value"`
	//Hash - хэш
	Hash string `db:"hash"`
	//Create - время создания
	Create time.Time `db:"create_at"`
}
