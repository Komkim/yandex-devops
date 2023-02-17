package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"yandex-devops/storage"
)

const DBTIMEOUT = 5

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
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()
	err := f.PGXpool.QueryRow(ctx,
		`with unic_name as (
					select distinct name as n
					from metrics
				)
				select id, name, type, value, delta, hash, create_at
				from metrics,
					 unic_name
				where id in (select id from metrics where name = unic_name.n order by id desc limit 1)
					and name=$1;`,
		key,
	).Scan(&metric.ID, &metric.Name, &metric.MType, &metric.Value, &metric.Delta, &metric.Hash, &metric.Create)

	if err != nil {
		return nil, err
	}

	return convert(metric), nil
}

func (f PostgreStorage) GetAll() ([]storage.Metrics, error) {
	var m []storage.Metrics
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()
	rows, err := f.PGXpool.Query(ctx,
		`with unic_name as (
					select distinct name as n
					from metrics
				)
				select id, name, type, value, delta, hash, create_at
				from metrics,
					 unic_name
				where id in (select id from metrics where name = unic_name.n order by id desc limit 1);`,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		metric := metrics{}
		err := rows.Scan(&metric.ID, &metric.Name, &metric.MType, &metric.Value, &metric.Delta, &metric.Hash, &metric.Create)
		if err != nil {
			return nil, err
		}
		m = append(m, *convert(metric))
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (f PostgreStorage) SetOne(m storage.Metrics) (*storage.Metrics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into metrics (name, type, value, delta, hash)
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
	return &m, nil
}

func (f PostgreStorage) SetAll(metrics []storage.Metrics) ([]storage.Metrics, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	valueArgs := [][]interface{}{}
	tx, err := f.PGXpool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	if len(metrics) < 1 {
		return nil, nil
	}

	for _, m := range metrics {
		valueArgs = append(valueArgs, []interface{}{m.ID, m.MType, m.Value, m.Delta, m.Hash})
	}

	_, err = tx.CopyFrom(
		ctx,
		pgx.Identifier{"metrics"},
		[]string{"name", "type", "value", "delta", "hash"},
		pgx.CopyFromRows(valueArgs),
	)

	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	return metrics, nil
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
		ID:    m.Name,
		MType: m.MType,
		Delta: delta,
		Value: value,
		Hash:  m.Hash,
	}
}
