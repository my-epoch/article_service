package config

type Config struct {
	GrpcServerConfig GrpcServerConfig `yaml:"grpc"`
}

type GrpcServerConfig struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}
