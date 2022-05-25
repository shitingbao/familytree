package config

type Mysql struct {
	Host     string `toml:"host"`
	Port     uint32 `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	LogMode  int    `toml:"log_mode"`
}

type Mongo struct {
	Host     string `toml:"host"`
	Port     uint32 `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}
type Redis struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type Postgresql struct {
	Host string `toml:"host"`
}
