package infra

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// BootMysql return instance DB of gorm
func BootMysql() {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	password := viper.GetString("database.pass")
	username := viper.GetString("database.user")
	database := viper.GetString("database.name")

	domain := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", username, password, host, port, database)
	db, err := gorm.Open("mysql", domain)
	if err != nil {
		log.Panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Panic(err)
	}

	db.LogMode(!IsProduction())
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(30)
	db.DB().SetConnMaxLifetime(time.Second * 60)

	DB = db
}
