package config

type GRPCConfig struct {
	Auth    *GRPCServer
	OpenApi *GRPCOpenapi
}

type GRPCOpenapi struct {
	Mgcc     *GRPCServer
	Ocean    *GRPCServer
	KuaiShou *GRPCServer
}

type GRPCServer struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}
