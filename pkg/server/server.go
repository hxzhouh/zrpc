package server

import (
	"fmt"
	"github.com/hxzhouh/zrpc/pkg/core"
	etcd2 "github.com/hxzhouh/zrpc/pkg/discovery/etcd"
	"github.com/hxzhouh/zrpc/pkg/logger"
	"github.com/hxzhouh/zrpc/pkg/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	//etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/naming"
	"net"
)

type Options struct {
	Name    string
	Host    string
	Version string
	lis     net.Listener
	server  *grpc.Server
}

type Server struct {
	Options
	register *etcd2.Register
	config   core.ServiceConfig
}

func (s *Server) Init() {
	ip, err := utils.GetClientIp()
	if err != nil {
		logger.DefaultLogger.Fatal("ge local ip error")
	}
	ip = "0.0.0.0"
	s.config = *core.Config
	lis, err := net.Listen("tcp", ":0") //开启监听.:0 会随机分配端口
	port := lis.Addr().(*net.TCPAddr).Port
	s.Host = fmt.Sprintf("%s:%d", ip, port)
	logger.DefaultLogger.Info("listening on:", zap.String("host:port", s.Host)) // 地址(含端口)
	if err != nil {
		logger.DefaultLogger.Fatal(err.Error())
	}
	s.lis = lis
	s.server = grpc.NewServer() //新建一个grpc服务
	s.register = etcd2.NewRegister(s.config.Etcd, logger.DefaultLogger)

	go func() {
		timer := time.NewTicker(10 * time.Second)
		for {
			select {
			case <-timer.C:
				ss, _ := s.register.GetServerInfo()
				logger.DefaultLogger.Info(ss.Addr)
			}
		}
	}()
}

func (s *Server) Start() {
	go func() {
		if err := s.server.Serve(s.lis); err != nil {
			logger.DefaultLogger.Fatal("failed to serve:", zap.Error(err))
		}
	}()
}
func (s *Server) Run() {
	s.Start() // 启动服务，
	// do something
	_, err := s.register.Register(etcd2.Server{ // 注册到服务发现。
		Name:    s.Name,
		Addr:    s.Host,
		Version: s.Version,
		Weight:  10,
	}, 10)
	if err != nil {
		logger.DefaultLogger.Fatal("register error", zap.Error(err))
	}
	c := make(chan os.Signal, 1) // 优雅停机
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		t := <-c
		switch t {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logger.DefaultLogger.Info("service stop...", zap.String("service Name", s.Name))
			s.register.Stop() // 从服务注册摘除
			s.server.Stop()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func (s *Server) Server() *grpc.Server {
	return s.server
}

func NewServer(name, version string) *Server {
	return &Server{Options: Options{
		Name:    name,
		Version: version,
	}}
}
