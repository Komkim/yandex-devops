package postgresql

import (
	"context"
	"github.com/stretchr/testify/require"
)

const skipTestMessage = "Skip test. please up local database for this test"

func getStorageRepo(req *require.Assertions) PostgreStorage {
	ctx := context.Background()

	//type Config struct {
	//	DatabaseDSN string `env:"DATABASE_URI" mapstructure:"databasedsn"`
	//}
	//type cfg struct {
	//	Dsn *Config
	//}
	//
	//var c cfg
	//err := envconfig.Init(&c)
	//req.NoError(err)
	//repo, err := New(ctx, c.Dsn.DatabaseDSN)
	dsn := "postgres://postgres:changeme@localhost:5432/yandex?sslmode=disable"
	repo, err := New(ctx, dsn)

	req.NoError(err)

	return *repo
}
