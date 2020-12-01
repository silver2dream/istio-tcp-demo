package database

import (
	"errors"
	"fmt"
	"server/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func initDbConn(config *conf.Database) error {
	conninfo := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true", config.User, config.Passwd, config.Host, config.Db)
	var err error
	Instance, err = gorm.Open(mysql.Open(conninfo), &gorm.Config{}) //gorm.Open("mysql", conninfo)
	if err != nil {
		msg := fmt.Sprintf("Failed to connect to db '%s', err: %s", conninfo, err.Error())
		return errors.New(msg)
	}

	return nil
}
