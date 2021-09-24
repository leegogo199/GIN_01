package Controller

import (
	"GINVUE/Model"
	"GINVUE/common"
	"GINVUE/dto"
	"GINVUE/response"
	"GINVUE/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	db := common.GetDB()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity,
			422, nil, "手机号必须为11位")

		return
	}
	if len(password) < 6 || len(password) > 11 {
		response.Response(ctx, http.StatusUnprocessableEntity,
			422, nil, "密码必须大于6位且小于11位")
		return
	}
	//
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, telephone, password)
	//判断手机号是否存在
	if util.IsTelephoneExist(db, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity,
			422, nil, "该用户已经注册")
		return
	}
	//密码加密
	hashPD, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError,
			500, nil, "加密错误")
		return
	}
	//创建用户

	newUser := Model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPD),
	}
	db.Create(&newUser)

	//返回结果
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "注册成功",
	})

	return
}

func Login(ctx *gin.Context) {
	//获取参数
	DB := common.GetDB()
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 || len(password) > 11 {
		ctx.JSON(http.StatusUnprocessableEntity,
			gin.H{"code": 422, "msg": "密码必须大于6位且小于11位"})
		return
	}
	//手机号判断
	var user Model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "用户不存在"})
		return
	}

	//判断密码正确？
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}

	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登陆成功",
	})
	response.Success(ctx, gin.H{"token": token}, "注册成功")

}
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(Model.User))},
	})

}
