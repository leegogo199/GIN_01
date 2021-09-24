package main

import (
	"GINVUE/Controller"
	"GINVUE/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", Controller.Register)
	r.POST("/api/auth/login", Controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), Controller.Info)
	return r
}
