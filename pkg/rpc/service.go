package rpc

import "fmt"

type Config struct {
	Name     string `yaml:"name"`
	Bind     string `yaml:"bind"`
	HttpBind string `yaml:"http_bind"`
}

func GetSvc(conf *Config) *Service {
	return &Service{
		Name: conf.Name,
		Bind: conf.Bind,
	}
}

func GetKey(svc *Service) string {
	return fmt.Sprintf("/%s/%s", svc.Name, svc.Bind)
}

func GetWatcherKey(name string) string {
	return fmt.Sprintf("/%s/", name)
}
