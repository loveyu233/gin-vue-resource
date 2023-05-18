package router

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

func RegisterRotationMap(router *gin.RouterGroup) {
	rm := router.Group("/rotationMap")
	rm.GET("/get", api.GetRotationMap)
}
