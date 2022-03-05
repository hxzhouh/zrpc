package core

import (
	"flag"
	"github.com/hxzhouh/zrpc/pkg/logger"
	"go.uber.org/zap"
	"strings"
)

func init() {
	Config = new(ServiceConfig)
	var etcd string
	flag.StringVar(&etcd, "etcd.addr", "10.252.176.88:2381,10.252.176.90:2381,10.252.176.89:2381", "etcd 地址")
	flag.Parse()
	if etcd == "" {
		logger.DefaultLogger.Fatal("Etcd.addr is nil")
	}
	Config.Etcd = strings.Split(etcd, ",")
	logger.DefaultLogger.Info("the config is ", zap.Any("config", *Config))
}
