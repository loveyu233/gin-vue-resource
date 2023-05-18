package api

import (
	"github.com/gin-gonic/gin"
	"server/model/response"
	"server/service"
)

func GetRotationMap(c *gin.Context) {
	r := service.RotationMap()
	response.ResOk(c, r)
}
