package fixtures

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Fixture interface {
	GetMetricsSQL() []string
}

type CleanupFixture struct{}

func (cf CleanupFixture) GetMetricsSQL() []string {
	return []string{
		`TRUNCATE TABLE metrics RESTART IDENTITY CASCADE;`,
	}
}

func ExecuteFixture(ctx context.Context, db *pgxpool.Pool, fixture Fixture) {
	for _, query := range fixture.GetMetricsSQL() {
		_, err := db.Exec(ctx, query)

		if err != nil {
			panic(err)
		}
	}
}
