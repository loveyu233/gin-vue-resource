package dao

import (
	"gorm.io/gorm"
	"strconv"
	"time"
)

type User struct {
	Uid              int    `json:"uid" gorm:"column:uid"`                             // id
	Username         string `json:"username" gorm:"column:username"`                   // 名称
	Password         string `json:"password" gorm:"column:password"`                   // 密码
	Email            string `json:"email" gorm:"column:email"`                         // 邮箱
	Phone            string `json:"phone" gorm:"column:phone"`                         // 手机号
	BackgroundImg    string `json:"background_img" gorm:"column:background_img"`       // 背景图
	SculptureHead    string `json:"sculpture_head" gorm:"column:sculpture_head"`       // 头像
	IntroduceBriefly string `json:"introduce_briefly" gorm:"column:introduce_briefly"` // 个人简介
	NumberCoin       int    `json:"number_coin" gorm:"column:number_coin"`             // 硬币数量
	NumberBCoin      int    `json:"number_b_coin" gorm:"column:number_b_coin"`         // 硬币数量
	Experience       int    `json:"experience" gorm:"column:experience"`               // 当前经验值
	NumberFans       int    `json:"number_fans" gorm:"column:number_fans"`             // 粉丝数量
	NumberFollow     int    `json:"number_follow" gorm:"column:number_follow"`         // 关注博主的数量
	IsDelete         bool   `json:"is_delete" gorm:"column:is_delete"`                 // 是否注销
	CreateTime       string `json:"create_time" gorm:"column:create_time"`             // 创建时间
	DeleteTime       string `json:"delete_time" gorm:"column:delete_time"`             // 注销时间
}

func (m *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreateTime = strconv.FormatInt(time.Now().Unix(), 10)
	return nil
}
