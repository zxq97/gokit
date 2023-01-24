package mysql

import (
	"fmt"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

const (
	mysqlAddr = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	DB       string `yaml:"db"`
	Password string `yaml:"password"`
}

func NewMysqlSess(conf *Config) (sqlbuilder.Database, error) {
	url := fmt.Sprintf(mysqlAddr, conf.User, conf.Password, conf.Host, conf.Port, conf.DB)
	dsn, err := mysql.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return mysql.Open(dsn)
}
