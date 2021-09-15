package main

import (
	"GINVUE/Controller"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/registers", Controller.Register)
	return r
}
