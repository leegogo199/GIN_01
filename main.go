// main.go
package main

import (
	//"GINVUE/Controller"
	//"GINVUE/Model"
	"GINVUE/common"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jinzhu/gorm"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	panic(r.Run())

}
