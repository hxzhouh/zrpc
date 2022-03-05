package zgrpc

import (
	"context"
	"github.com/hxzhouh/zrpc/pkg/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

// Server ...
type Server struct {
	*grpc.Server
	listener net.Listener
	*Config
}

// 新建一个grpc 服务
func newServer(config *Config) (*Server, error) {
	newServer := grpc.NewServer(config.serverOptions...)
	config.logger.Info("listen:", zap.String("address", config.Address()))

	listener, err := net.Listen(config.Network, config.Address())
	if err != nil {
		config.logger.Error("net.Listen failed", zap.Error(err))
		return nil, err
	}
	config.Port = listener.Addr().(*net.TCPAddr).Port

	return &Server{
		Server:   newServer,
		listener: listener,
		Config:   config,
	}, nil
}

func (s *Server) Healthz() bool {
	conn, err := s.listener.Accept()
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// Server implements server.Server interface.
func (s *Server) Serve() error {
	err := s.Server.Serve(s.listener)
	return err
}

// Stop implements server.Server interface
// it will terminate echo server immediately
func (s *Server) Stop() error {
	s.Server.Stop()
	return nil
}

// GracefulStop implements server.Server interface
// it will stop echo server gracefully
func (s *Server) GracefulStop(ctx context.Context) error {
	s.Server.GracefulStop()
	return nil
}

// Info returns server info, used by governor and consumer balancer
func (s *Server) Info() *server.ServiceInfo {
	return nil
}
