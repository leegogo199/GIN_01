package Controller

import (
	"GINVUE/Model"
	"GINVUE/common"
	"GINVUE/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": 422, "msg": "密码必须大于6位"})
		return
	}
	//
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if IsTelephoneExist(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": 422, "msg": "该用户已经注册"})
		return
	}
	//创建用户
	newUser := Model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)

	//返回结果
	ctx.JSON(200, gin.H{
		"message": "注册成功",
	})
	return
}
func IsTelephoneExist(db *gorm.DB, telephone string) bool {
	var user Model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
