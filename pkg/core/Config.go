package core

var Config *ServiceConfig

type ServiceConfig struct {
	Etcd []string
}
