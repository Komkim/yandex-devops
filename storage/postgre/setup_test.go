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
		DSN string `env:"DSN_DATABASE" mapstructure:"databasedsn"`
	}
	type cfg struct {
		c *Config
	}

	var cc cfg
	err := envconfig.Init(&cc)
	req.NoError(err)

	repo, err := New(ctx, cc.c.DSN)

	req.NoError(err)

	return *repo
}
