package postgresql

import (
	"database/sql"
	"time"
)

type metrics struct {
	ID     string          `db:"id"`
	Name   string          `db:"name"`
	MType  string          `db:"type"`
	Delta  sql.NullInt64   `db:"delta"`
	Value  sql.NullFloat64 `db:"value"`
	Hash   string          `db:"hash"`
	Create time.Time       `db:"create_at"`
}
