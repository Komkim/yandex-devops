package postgresql

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/vrischmann/envconfig"
	"yandex-devops/config"
)

const skipTestMessage = "Skip test. please up local database for this test"

func getStorageRepo(req *require.Assertions) PostgreStorage {
	ctx := context.Background()

	//type Config struct {
	//	DSN string `env:"DSN_DATABASE" mapstructure:"databasedsn"`
	//}
	//type cfg struct {
	//	c *Config
	//}
	//
	//var cc cfg

	type cfg struct {
		Dsn *config.Server
	}

	var c cfg
	err := envconfig.Init(&c)
	req.NoError(err)

	repo, err := New(ctx, c.Dsn.DatabaseDSN)

	req.NoError(err)

	return *repo
}
