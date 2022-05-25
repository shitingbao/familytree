package conf

import (
	"flag"

	"familytree/family/lib/config"

	"github.com/BurntSushi/toml"
)

var (
	ConfigPath string
	Conf       = &Config{}
)

//TODO 后期所有的配置独立化 作为公共库配置
type Config struct {
	DB      *config.Postgresql `toml:"db"`
	HTTPSvc *config.HTTPServer `toml:"http_svc"`
}

func init() {
	flag.StringVar(&ConfigPath, "conf", "./config.toml", "default config path")
}

func Init() *Config {
	var (
		err error
	)

	_, err = toml.DecodeFile(ConfigPath, &Conf)
	if err != nil {
		panic(err)
	}

	return Conf
}
