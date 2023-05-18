package router

import "github.com/gin-gonic/gin"

func RegisterGlobalRouter(router *gin.RouterGroup) {
	RegisterUserRouter(router)
	RegisterUserStatic(router)
	RegisterRotationMap(router)
}
