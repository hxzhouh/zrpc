package client

import (
	"github.com/hxzhouh/zrpc/pkg/core"
	"github.com/hxzhouh/zrpc/pkg/discovery/etcd"
	"github.com/hxzhouh/zrpc/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

type Options struct {
	client      *grpc.ClientConn
	ServiceName string
}

type Client struct {
	Options
	config *core.ServiceConfig
}

func (c *Client) Init() {
	c.config = core.Config

}

func (c *Client) Run() (*grpc.ClientConn, error) {
	options := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	}
	r := etcd.NewResolver(c.config.Etcd, logger.DefaultLogger)
	resolver.Register(r)

	conn, err := grpc.Dial("etcd:///"+c.ServiceName, options...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
func NewClient(serverName string) *Client {
	return &Client{
		Options: Options{ServiceName: serverName},
	}
}
