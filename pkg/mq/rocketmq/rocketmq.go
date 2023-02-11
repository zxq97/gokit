package rocketmq

type Config struct {
	Addr   []string `yaml:"addr"`
	NsAddr []string `yaml:"ns_addr"`
}
