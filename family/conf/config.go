package conf

import (
	"flag"

	"familytree/family/lib/config"

	"github.com/BurntSushi/toml"
)

var (
	ConfigPath     string
	GRPCConfigPath string
	Conf           = &Config{}
)

//TODO 后期所有的配置独立化 作为公共库配置
type Config struct {
	DB      *config.Mysql      `toml:"db"`
	HTTPSvc *config.HTTPServer `toml:"http_svc"`
	GRPCSvc *config.GRPCConfig `toml:"grpc"`
	JOBSvc  *config.JobConfig  `toml:"job"`
	Qbox    *Qbox              `toml:"qbox"`
}

type Qbox struct {
	AccessKey   string `toml:"access_key"`
	SecretKey   string `toml:"secret_key"`
	Bucket      string `toml:"bucket"`
	VideoBucket string `toml:"video_bucket"`
}

func init() {
	flag.StringVar(&ConfigPath, "conf", "./config.toml", "default config path")
	flag.StringVar(&GRPCConfigPath, "grpc", "./grpc.toml", "default grpc config path")
}

func Init() *Config {

	var (
		err error
	)

	_, err = toml.DecodeFile(ConfigPath, &Conf)
	if err != nil {
		panic(err)
	}

	_, err = toml.DecodeFile(GRPCConfigPath, &Conf.GRPCSvc)
	if err != nil {
		panic(err)
	}

	return Conf
}
