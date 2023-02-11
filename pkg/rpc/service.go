package rpc

import (
	"fmt"

	"github.com/zxq97/gokit/internal/host"
)

type Config struct {
	Name     string `yaml:"name"`
	Bind     string `yaml:"bind"`
	HttpBind string `yaml:"http_bind"`
}

func GetSvc(conf *Config) (*Service, error) {
	ip, err := host.GetIP()
	return &Service{
		Name: conf.Name,
		Bind: ip + conf.Bind,
	}, err
}

func GetKey(svc *Service) string {
	return fmt.Sprintf("/%s/%s", svc.Name, svc.Bind)
}

func GetWatcherKey(name string) string {
	return fmt.Sprintf("/%s/", name)
}
