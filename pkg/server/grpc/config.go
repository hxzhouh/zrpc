package grpc

import (
	"github.com/hxzhouh/zrpc/pkg/flag"
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
		Network:                   "tcp4",
		Host:                      flag.String("host"),
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
