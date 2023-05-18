package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"math"
	"server/cache"
	"server/global"
	. "server/global"
	"server/model/dao"
	"server/model/dto"
	"server/model/response"
	"server/model/system"
	"server/util/imageVerificationCode"
	"server/util/jwt"
	"server/util/level"
	"server/util/random"
	"strconv"
	"time"
)

type UserService struct{}

// FindUID 根据uid返回用户数据
func (a UserService) FindUID(uid int) *dto.UserVo {
	var userVo dto.UserVo
	var userDao dao.User
	rowsAffected := GormClient.Where(&dao.User{Uid: uid}).First(&userDao).RowsAffected
	if rowsAffected == 0 {
		return nil
	}
	copier.Copy(&userVo, userDao)
	userVo.Level = level.CalculateLevel(userVo.Experience)
	userVo.ExperienceRequired = level.CalculateExperience(userVo.Level + 1)
	return &userVo
}

func (a UserService) Login(user dto.UserLogin) response.R {
	x := cache.Get(user.Email)
	x1, _ := strconv.ParseFloat(x, 10)
	x2, _ := strconv.ParseFloat(user.X, 10)
	abs := math.Abs(x1 - x2)
	if abs > 10 {
		return response.VerificationCodeVerificationError
	}
	var daoUser dao.User
	var count int64
	GormClient.Where(&dao.User{Email: user.Email, Password: user.Password}).First(&daoUser).Count(&count)
	if count < 1 {
		return response.UserDataDoesNotExist
	}
	token, err := jwt.GenerateAccessToken(daoUser.Uid)
	if err != nil {
		return response.TokenCreateError
	}
	var uservo dto.UserVo
	copier.Copy(&uservo, &daoUser)
	uservo.Level = level.CalculateLevel(uservo.Experience)
	uservo.ExperienceRequired = level.CalculateExperience(uservo.Level+1) - uservo.Experience
	return response.NewR(gin.H{"token": token, "user": uservo})
}

func (a UserService) GetVerCode(email string) response.R {
	cache.Set(email, random.GetCode(), time.Minute*3)
	return response.NewR("")
}

func (a UserService) Register(userCode dto.UserVerificationCode) response.R {
	if cache.Get(userCode.Email) != userCode.Code {
		return response.VerificationCodeVerificationError
	}
	var daoUser dao.User
	var uservo dto.UserVo
	if GormClient.Where(&dao.User{Email: userCode.Email}).First(&daoUser).RowsAffected == 0 {
		user := geoDaoUser(userCode.Email)
		if GormClient.Create(&user).RowsAffected == 0 {
			return response.CreateDataError
		}
		token, err := jwt.GenerateAccessToken(user.Uid)
		if err != nil {
			return response.TokenCreateError
		}
		copier.Copy(&uservo, user)
		return response.NewR(gin.H{"token": token})
	}
	token, err := jwt.GenerateAccessToken(daoUser.Uid)
	if err != nil {
		return response.TokenCreateError
	}
	copier.Copy(&uservo, daoUser)
	uservo.Level = level.CalculateLevel(uservo.Experience)
	uservo.ExperienceRequired = level.CalculateExperience(uservo.Level+1) - uservo.Experience
	return response.NewR(gin.H{"token": token, "user": uservo})
}

func geoDaoUser(email string) dao.User {
	return dao.User{
		Email:    email,
		Username: "gin-vue-resource",
	}
}

var ver *imageVerificationCode.Jigsaw

func init() {
	ver = imageVerificationCode.New()
	ver.SetBgDir("./static/jigsaw/bg/")
	ver.SetMaskPath("./static/jigsaw/mask.png")
}

func (a UserService) GetCode(email string) response.R {
	verDao, err := imageVerificationCode.Create()
	if err != nil {
		return response.VerificationCodeCreateError
	}
	var verVo system.VerificationCodeVo
	copier.Copy(&verVo, &verDao)
	cache.Set(email, verDao.X, global.VerCodeTime)
	return response.NewR(verVo)
}
