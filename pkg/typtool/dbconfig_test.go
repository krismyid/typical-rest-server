package typtool_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-rest-server/pkg/typtool"
)

func TestDBConfig(t *testing.T) {
	os.Setenv("ABC_HOST", "some-host")
	os.Setenv("ABC_PORT", "some-port")
	os.Setenv("ABC_DBNAME", "some-dbname")
	os.Setenv("ABC_DBUSER", "some-dbuser")
	os.Setenv("ABC_DBPASS", "some-dbpass")
	defer os.Clearenv()
	redisConfig := typtool.DBEnvKeysWithPrefix("ABC")
	require.Equal(t, &typtool.DBConfig{
		Host:   "some-host",
		Port:   "some-port",
		DBName: "some-dbname",
		DBUser: "some-dbuser",
		DBPass: "some-dbpass",
	}, redisConfig.Config())
}
