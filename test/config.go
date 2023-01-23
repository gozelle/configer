package test

type Config struct {
	Addr     *string  `toml:"addr" default:""`
	Database *string  `toml:"database"`
	Monitor  *Monitor `toml:"monitor"`
}

type Database struct {
	Host *string `toml:"host" default:"127.0.0.1"`
}

type Auth struct {
	Username string `toml:"username" default:"test"`
	Password string
}

type Monitor struct {
	Addr *string `toml:"addr"`
}
