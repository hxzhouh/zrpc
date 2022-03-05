package zgrpc

import (
	"fmt"
	"github.com/hxzhouh/zrpc/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Config struct {
	Name       string `json:"name"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Deployment string `json:"deployment"`
	// Network network type, tcp4 by default
	Network string `json:"network" toml:"network"`
	// EnableAccessLog enable Access Interceptor, true by default
	EnableAccessLog bool
	// DisableTrace disbale Trace Interceptor, false by default
	DisableTrace bool
	// DisableMetric disable Metric Interceptor, false by default
	DisableMetric bool
	// SlowQueryThresholdInMilli, request will be colored if cost over this threshold value
	SlowQueryThresholdInMilli int64
	// ServiceAddress service address in registry info, default to 'Host:Port'
	ServiceAddress string
	// EnableTLS
	EnableTLS bool
	// CaFile
	CaFile string
	// CertFile
	CertFile string
	// PrivateFile
	PrivateFile string

	Labels map[string]string `json:"labels"`

	serverOptions      []grpc.ServerOption
	streamInterceptors []grpc.StreamServerInterceptor
	unaryInterceptors  []grpc.UnaryServerInterceptor

	logger *zap.Logger
}

func DefaultConfig() *Config {
	return &Config{
		Network:                   "tcp",
		Host:                      "",
		Port:                      9092,
		Deployment:                "",
		EnableAccessLog:           true,
		DisableMetric:             false,
		DisableTrace:              false,
		EnableTLS:                 false,
		SlowQueryThresholdInMilli: 500,
		logger:                    logger.DefaultLogger,
		serverOptions:             []grpc.ServerOption{},
		streamInterceptors:        []grpc.StreamServerInterceptor{},
		unaryInterceptors:         []grpc.UnaryServerInterceptor{},
	}
}

func StdConfig(name string) *Config {
	return RawConfig("jupiter.server." + name)
}

// RawConfig ...
func RawConfig(key string) *Config {
	var config = DefaultConfig()
	return config
}

// 从配置中创建grpc。
func (config *Config) MustBuild() *Server {
	server, err := config.Build()
	if err != nil {
		logger.DefaultLogger.Fatal("build xgrpc server: %v", zap.Error(err))
	}
	return server
}

// Build ...
func (config *Config) Build() (*Server, error) {
	//if !config.DisableTrace {
	//	config.unaryInterceptors = append(config.unaryInterceptors, traceUnaryServerInterceptor)
	//	config.streamInterceptors = append(config.streamInterceptors, traceStreamServerInterceptor)
	//}
	//
	//if !config.DisableMetric {
	//	config.unaryInterceptors = append(config.unaryInterceptors, prometheusUnaryServerInterceptor)
	//	config.streamInterceptors = append(config.streamInterceptors, prometheusStreamServerInterceptor)
	//}

	return newServer(config)
}

// Address ...
func (config Config) Address() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}
