package main

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Info struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Role         string  `gorm:"size:255;default:'小王子'"`
	Num          string `gorm:"default:'0'"`
}

func main() {
	db, err := gorm.Open("mysql", "root:940213@tcp(49.234.156.231:3306)/kkb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&Info{})

	db.Debug()
}
