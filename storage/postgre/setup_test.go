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
		DSN string
	}
	type cfg struct {
		c *Config
	}

	var c cfg

	//type cfg struct {
	//	Dsn *config.Server
	//}

	//var c cfg
	err := envconfig.Init(&c)
	req.NoError(err)

	repo, err := New(ctx, c.c.DSN)

	req.NoError(err)

	return *repo
}
