package main

import (
	"server/util/jwt"
	"server/util/level"
	"testing"
)

func TestFunc(t *testing.T) {
	// daoUser copy to dtoUser
	//copier.Copy(dtoUser, &daoUser)

	// 验证
	//validator.New().Struct(stu)
}

func TestA(t *testing.T) {
	token, _ := jwt.GenerateAccessToken(1)
	t.Log(token)
	//signature is invalid
	_, claims, err := jwt.ParseToken(token + "a")
	t.Log(claims.UserId, err)
}

func TestImg(t *testing.T) {
	t.Log(level.CalculateLevel(200))
	t.Log(level.CalculateExperience(3) - 200)
}
