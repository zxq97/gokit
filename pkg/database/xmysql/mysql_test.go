package xmysql

import (
	"log"
	"testing"

	"github.com/zxq97/gokit/pkg/config"
)

func TestNewMysqlDB(t *testing.T) {
	conf := &Config{}
	err := config.LoadYaml("../../../internal/yaml/xmysql.yaml", conf)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(conf)
	db, err := NewMysqlDB(conf)
	if err != nil {
		t.Fatal(err)
	}
	type User struct {
		UID      int64
		Nickname string
	}
	u := &User{}
	err = db.Model(&User{}).First(u).Error
	if err != nil {
		t.Fatal(err)
	}
	log.Println(u)
	u.UID = 99999
	err = db.Create(u).Error
	if err != nil {
		t.Fatal(err)
	}
}
