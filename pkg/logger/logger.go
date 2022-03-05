package logger

import "go.uber.org/zap"

var DefaultLogger *zap.Logger

// init log
func init() {
	DefaultLogger, _ = zap.NewProduction() // todo 后续完善zap配置。
}
