package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
	"time"
	"yandex-devops/storage"
)

type PostgreStorage struct {
	PGXpool *pgxpool.Pool
}

func New(ctx context.Context, connString string) (*PostgreStorage, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}
	return &PostgreStorage{pool}, nil
}

func (f PostgreStorage) GetOne(key string) (*storage.Metrics, error) {
	metric := metrics{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := f.PGXpool.QueryRow(ctx, "select * from metrics where id = ?", key).Scan(metric)
	if err != nil {
		return nil, err
	}

	return convert(metric), nil
}

func (f PostgreStorage) GetAll() (*[]storage.Metrics, error) {
	var m []storage.Metrics
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rows, err := f.PGXpool.Query(ctx, "select * from metrics")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		metric := metrics{}
		err := rows.Scan(metric)
		if err != nil {
			return nil, err
		}
		m = append(m, *convert(metric))
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (f PostgreStorage) SetOne(metric storage.Metrics) (*storage.Metrics, error) {
	var m storage.Metrics
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sqlStatement := `
		insert into mertics (name, type, value, delta, hash)
		values ($1, $2, $3, $4, $5)
		returning id `
	var id int
	err := f.PGXpool.QueryRow(ctx, sqlStatement, m.ID, m.MType, m.Value, m.Delta, m.Hash).Scan(&id)
	if err != nil {
		return nil, err
	}

	if id < 1 {
		return nil, nil
	}
	return &metric, nil
}

func (f PostgreStorage) SetAll(metrics []storage.Metrics) (*[]storage.Metrics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	valueStr := []string{}
	valueArgs := []interface{}{}
	tx, err := f.PGXpool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	if len(metrics) < 1 {
		return nil, nil
	}

	for _, m := range metrics {
		valueStr = append(valueStr, "?, ?, ?, ?, ?, ?")

		valueArgs = append(valueArgs, m.ID, m.MType, m.Value, m.Delta, m.Hash, time.Now())
	}

	sqlStatement := `insert into mertics (name, type, value, delta, hash) values`
	sqlStatement = fmt.Sprintf(sqlStatement, strings.Join(valueStr, ","))

	cTag, err := tx.Exec(ctx, sqlStatement, valueArgs)
	//_ , err = tx.Exec(ctx, sqlStatement)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	if cTag.RowsAffected() != int64(len(metrics)) {
		tx.Rollback(ctx)
		return nil, errors.New("rows affected error")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return &metrics, nil
}

func (f PostgreStorage) Close() error {
	f.PGXpool.Close()
	return nil
}

func convert(m metrics) *storage.Metrics {
	var value *float64
	var delta *int64
	if m.Value.Valid {
		value = &m.Value.Float64
	}
	if m.Delta.Valid {
		delta = &m.Delta.Int64
	}
	return &storage.Metrics{
		ID:    m.ID,
		MType: m.MType,
		Delta: delta,
		Value: value,
		Hash:  m.Hash,
	}
}
