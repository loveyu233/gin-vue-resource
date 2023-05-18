package api

import (
	"github.com/gin-gonic/gin"
	"server/model/dto"
	. "server/model/response"
	"server/service"
	"server/util/validator"
	"server/util/verify"
)

var userService = new(service.UserService)

// 登录
func Login(c *gin.Context) {
	var dtoUser dto.UserLogin
	if err := c.ShouldBind(&dtoUser); err != nil {
		ResOk(c, validator.ValidatorTranslatorErr(err))
		return
	}
	r := userService.Login(dtoUser)
	ResOk(c, r)
}

// 获取验证码
func GetNumberCode(c *gin.Context) {
	email := c.Query("email")
	if !verify.VerifyEmailFormat(email) {
		ResOk(c, ParameterValidationError)
		return
	}
	r := userService.GetVerCode(email)
	ResOk(c, r)
}

// 注册/验证码登录
func Register(c *gin.Context) {
	user := dto.UserVerificationCode{}
	err := c.ShouldBind(&user)
	if err != nil {
		r := validator.ValidatorTranslatorErr(err)
		ResOk(c, r)
		return
	}
	r := userService.Register(user)
	ResOk(c, r)
}

// 获取图形验证码
func GetImageCode(c *gin.Context) {
	email := c.Query("email")
	if !verify.VerifyEmailFormat(email) {
		ResOk(c, ParameterValidationError)
		return
	}
	r := userService.GetCode(email)
	ResOk(c, r)
}

func Token(c *gin.Context) {

}
