package configer

import (
	"github.com/gozelle/testify/require"
	"testing"
)

type Config struct {
	Addr     *string  `toml:"addr" default:""`
	Database *string  `toml:"database"`
	Monitor  *Monitor `toml:"monitor"`
	Auth     *Auth    `toml:"auth"`
}

type Database struct {
	Host *string `toml:"host" default:"127.0.0.1"`
	Auth *Auth   `toml:"auth"`
}

type Auth struct {
	Username string `toml:"username" default:"test"`
	Password string `toml:"password"`
}

type Monitor struct {
	*Auth
	Addr *string `toml:"addr"`
}

func TestLoad(t *testing.T) {
	addr := "127.0.0.1"
	conf := &Config{
		Addr: &addr,
	}
	err := handle(conf)
	require.NoError(t, err)
}
