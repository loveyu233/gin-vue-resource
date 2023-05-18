package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUserStatic(router *gin.RouterGroup) {
	static := router.Group("/static")
	static.StaticFS("/images", http.Dir("./static/images"))
	static.StaticFS("/video", http.Dir("./static/video"))
}
