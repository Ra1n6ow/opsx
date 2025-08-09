package center

import (
	"time"

	"github.com/ra1n6ow/opsx/internal/pkg/log"
	"github.com/spf13/viper"
)

// Config 运行时配置.
type Config struct {
	ServerMode string
	JWTKey     string
	Expiration time.Duration
}

// UnionServer 定义一个联合服务器. 根据 ServerMode 决定要启动的服务器类型.
type UnionServer struct {
	cfg *Config
}

// NewUnionServer 根据配置创建联合服务器(http,grpc,grpc-gateway)
func (cfg *Config) NewUnionServer() (*UnionServer, error) {
	return &UnionServer{cfg: cfg}, nil
}

// Run 运行应用.
func (s *UnionServer) Run() error {
	log.Infow("ServerMode from ServerOptions", "jwt-key", s.cfg.JWTKey)
	log.Infow("ServerMode from Viper", "jwt-key", viper.GetString("jwt-key"))

	select {}
}
