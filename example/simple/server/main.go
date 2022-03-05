package main

import (
	"context"
	"fmt"
	"github.com/hxzhouh/zrpc/example/simple/hello"
	"github.com/hxzhouh/zrpc/pkg/application"
	"github.com/hxzhouh/zrpc/pkg/logger"
	"github.com/hxzhouh/zrpc/pkg/server/zgrpc"
	"go.uber.org/zap"
)

type Engine struct {
	application.Application
}

type HelloServiceImpl struct {
}

func (eng *Engine) serveGRPC() error {
	server, err := zgrpc.StdConfig("zgrpc").Build()
	if err != nil {
		return err
	}
	hello.RegisterHelloServer(server.Server, new(HelloServiceImpl))
	return eng.Serve(server)
}

func (t *HelloServiceImpl) SayHelloStream(ctx context.Context, in *hello.HelloReq) (*hello.HelloResp, error) {
	return &hello.HelloResp{Name: fmt.Sprintf("hello,%s", in.Name)}, nil
}

func NewEngine() *Engine {
	eng := &Engine{}
	if err := eng.Startup(
		eng.serveGRPC,
	); err != nil {
		logger.DefaultLogger.Fatal("panic:", zap.Error(err))
	}
	return eng
}

func main() {
	eng := NewEngine()
	if err := eng.Run(); err != nil {
		logger.DefaultLogger.Error(err.Error())
	}
}
