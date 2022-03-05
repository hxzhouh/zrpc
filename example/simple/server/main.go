package main

import (
	"context"
	"fmt"
	"github.com/hxzhouh/zrpc/example/simple/hello"
	"github.com/hxzhouh/zrpc/pkg/application"
	"github.com/hxzhouh/zrpc/pkg/logger"
	"github.com/hxzhouh/zrpc/pkg/server"
	"google.golang.org/grpc"
)

type Engine struct {
	application.Application
}

func NewHelloService() *Engine {
	hello := &Engine{}
	if ok := hello.Startup(hello.serveGRPC()); ok != nil {

	}

	return hello
}

type HelloServiceImpl struct {
}

func (eng *Engine) serveGRPC() error {
	//server := xgrpc.StdConfig("grpc").Build()
	server := grpc.NewServer()
	hello.RegisterHelloServer(server, new(HelloServiceImpl))
	return eng.Serve(server)
}

func (t *HelloServiceImpl) SayHelloStream(ctx context.Context, in *hello.HelloReq) (*hello.HelloResp, error) {
	return &hello.HelloResp{Name: fmt.Sprintf("hello,%s", in.Name)}, nil
}

func main() {
	logger.DefaultLogger.Info("service running.....")
	s := server.NewServer("hello", "1.0.0")
	s.Init()
	hello.RegisterHelloServer(s.Server(), &HelloServiceImpl{})
	s.Run() // block here
}
