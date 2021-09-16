package main

import (
	"GINVUE/Controller"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", Controller.Register)
	r.POST("/api/auth/login", Controller.Login)
	return r
}
