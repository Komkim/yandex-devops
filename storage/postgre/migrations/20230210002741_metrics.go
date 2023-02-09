package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upMetrics, downMetrics)
}

func upMetrics(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table metrics
			(
				id        serial primary key,
				create_at timestamp with time zone default current_timestamp,
				name      varchar(40) not null,
				value     double precision,
				delta     integer,
				hash      varchar(100)
			);
	`)
	return err
}

func downMetrics(tx *sql.Tx) error {
	_, err := tx.Exec("drop table metrics")
	return err
}
