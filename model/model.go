package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Conf struct {
	Host     string `json:"host"`
	Port     string `json:"port'`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var gdb *gorm.DB

func NewDb(conf *Conf) error {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&autocommit=true",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err := gorm.Open("mysql", dbUri)
	if err != nil {
		return err
	}
	gdb = db
	return nil
}

func Close() {
	if gdb != nil {
		defer gdb.Close()
	}
}
