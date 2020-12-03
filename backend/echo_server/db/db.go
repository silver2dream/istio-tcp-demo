package database

import (
	"errors"
	"fmt"
	"server/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *gorm.DB

type Player struct {
	gorm.Model
	Name     string
	Account  string
	Password string
}

func InitDbConn(config conf.Database) error {
	if !config.External {
		return nil
	}

	conninfo := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true", config.User, config.Passwd, config.Host, config.Db)
	var err error
	instance, err = gorm.Open(mysql.Open(conninfo), &gorm.Config{}) //gorm.Open("mysql", conninfo)
	if err != nil {
		msg := fmt.Sprintf("Failed to connect to db '%s', err: %s", conninfo, err.Error())
		return errors.New(msg)
	}

	return nil
}

func GetTestingSQLService() string {
	var player Player
	player.Name = "Guest"
	if instance != nil {
		instance.First(&player, "account = ?", "test123")
	}
	return player.Name
}
