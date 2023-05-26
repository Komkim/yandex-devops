package postgresql

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/vrischmann/envconfig"
)

const skipTestMessage = "Skip test. please up local database for this test"

func getStorageRepo(req *require.Assertions) PostgreStorage {
	ctx := context.Background()

	type Config struct {
		DatabaseDSN string `env:"DSN_DATABASEDSN" mapstructure:"databasedsn"`
	}
	type cfg struct {
		Dsn *Config
	}

	var c cfg
	err := envconfig.Init(&c)
	req.NoError(err)

	repo, err := New(ctx, c.Dsn.DatabaseDSN)
	req.NoError(err)

	return *repo
}
