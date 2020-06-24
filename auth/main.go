package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"

	"g1/auth/common"
	"g1/auth/routers"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	routers.Routes(r)
	r.Run(":" + viper.GetString("server.port"))
}

func InitConfig()  {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/auth/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}