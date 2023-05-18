package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

func RegisterUserRouter(router *gin.RouterGroup) {
	user := router.Group("/user")
	user.POST("/login", api.Login)
	user.GET("/numbercode", api.GetNumberCode)
	user.GET("/imagecode", api.GetImageCode)
	user.POST("/register", api.Register)
	user.POST("/token", api.Token)
}
