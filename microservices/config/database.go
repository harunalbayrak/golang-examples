package config

import (
	"examples/microservices/pkg/setting"
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func DBUrl() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.AppSettings.DBSettings.Username,
		setting.AppSettings.DBSettings.Password,
		setting.AppSettings.DBSettings.Host,
		setting.AppSettings.DBSettings.Port,
		setting.AppSettings.DBSettings.DBName,
	)
}

func GetDB() *gorm.DB {
	return DB
}
