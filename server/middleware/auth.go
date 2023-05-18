package middleware

import (
	"github.com/gin-gonic/gin"
	"server/model/response"
	"server/util/jwt"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("token")
	_, claims, err := jwt.ParseToken(token)
	if err != nil {
		response.ResOk(c, response.TokenAuthorizationError)
		return
	}
	c.Set("uid", claims.UserId)
}
