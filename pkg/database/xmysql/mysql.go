package xmysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

const (
	mysqlAddr = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True"
)

type Config struct {
	Master   string   `yaml:"master"`
	Slave    []string `yaml:"slave"`
	Port     int      `yaml:"port"`
	User     string   `yaml:"user"`
	DB       string   `yaml:"db"`
	Password string   `yaml:"password"`
}

func NewMysqlDB(conf *Config) (*gorm.DB, error) {
	master := fmt.Sprintf(mysqlAddr, conf.User, conf.Password, conf.Master, conf.Port, conf.DB)
	slaves := make([]gorm.Dialector, len(conf.Slave))
	for k, v := range conf.Slave {
		dsn := fmt.Sprintf(mysqlAddr, conf.User, conf.Password, v, conf.Port, conf.DB)
		slaves[k] = mysql.Open(dsn)
	}
	db, err := gorm.Open(mysql.Open(master), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(master)},
		Replicas: slaves,
		Policy:   dbresolver.RandomPolicy{},
	}))
	if err != nil {
		return nil, err
	}
	return db, nil
}
