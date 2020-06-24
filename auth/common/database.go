package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	"g1/auth/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	driver := viper.GetString("datasource.driver")
	user := viper.GetString("datasource.user")
	password := viper.GetString("datasource.password")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, database,
		charset)
	db, err := gorm.Open(driver, args)
	if err != nil {
		panic(err)
	} else {
		db.AutoMigrate(&model.User{})
		DB = db
	}
	return db
}

func GetDB() *gorm.DB {
	return DB
}
