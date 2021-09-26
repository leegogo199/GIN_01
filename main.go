// main.go
package main

import (
	//"GINVUE/Controller"
	//"GINVUE/Model"
	"GINVUE/common"
	//"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	//"github.com/jinzhu/gorm"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)

	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}
func InitConfig() {
	WorkDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(WorkDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
