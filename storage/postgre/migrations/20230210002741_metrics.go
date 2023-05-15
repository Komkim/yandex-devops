// Модуль миграций
package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

// init - инициализация методов миграций
func init() {
	goose.AddMigration(upMetrics, downMetrics)
}

// upMetrics - запуск миграции
func upMetrics(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table metrics
			(
				id        serial primary key,
				create_at timestamp with time zone default current_timestamp,
				name      varchar(40) not null,
				type 	  varchar(40) not null,
				value     double precision,
				delta     bigint,
				hash      varchar(100)
			);
	`)
	return err
}

// downMetrics - откат миграции
func downMetrics(tx *sql.Tx) error {
	_, err := tx.Exec("drop table metrics")
	return err
}
