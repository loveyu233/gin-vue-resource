package service

import (
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model/dao"
	"server/model/response"
)

func RotationMap() response.R {
	var rms []dao.RotationMap
	global.GormClient.Model(&dao.RotationMap{}).Find(&rms)
	return response.NewR(gin.H{
		"rotationMap": rms,
	})
}
