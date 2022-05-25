package config

type HTTPServer struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}
