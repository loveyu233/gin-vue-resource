package core

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"server/global"
	"server/initialization"
	"server/middleware"
	"server/router"
	"server/util/validator"
)

func InitServer() {
	// 初始化
	{
		initialization.InitViper()
		initialization.InitLogger()
		global.GormClient = initialization.InitGorm()
		global.RedisClient = initialization.InitRedis()
		global.MongoClient = initialization.InitMongoDb()
	}
	validator.InitTrans("zh")
	engine := gin.Default()
	// 解决跨域,配置日志,重置错误
	engine.Use(middleware.Cors(), middleware.GinLogger, middleware.GinRecovery(false))
	routerGroup := engine.RouterGroup.Group("/api")
	router.RegisterGlobalRouter(routerGroup)
	err := engine.Run("127.0.0.1:" + viper.GetString("other.port"))
	if err != nil {
		panic(err.Error())
	}
}
