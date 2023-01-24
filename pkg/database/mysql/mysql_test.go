package mysql

import (
	"log"
	"testing"

	"github.com/zxq97/gokit/pkg/config"
)

func TestNewMysqlSess(t *testing.T) {
	conf := &Config{}
	err := config.LoadYaml("../../../internal/yaml/mysql.yaml", conf)
	if err != nil {
		t.Error()
	}
	log.Println(conf)
	sess, err := NewMysqlSess(conf)
	if err != nil {
		t.Error(err)
	}
	sql := "select * from users limit 1"
	row, err := sess.QueryRow(sql)
	if err != nil {
		t.Error(err)
	}
	log.Println(row)
}
