package vo

type UserShow struct {
	Uid              int    `json:"uid" gorm:"column:uid"`                             // id
	Username         string `json:"username" gorm:"column:username"`                   // 名称
	BackgroundImg    string `json:"background_img" gorm:"column:background_img"`       // 背景图
	SculptureHead    string `json:"sculpture_head" gorm:"column:sculpture_head"`       // 头像
	IntroduceBriefly string `json:"introduce_briefly" gorm:"column:introduce_briefly"` // 个人简介
	NumberCoin       int    `json:"number_coin" gorm:"column:number_coin"`             // 硬币数量
	Level            int    `json:"level" gorm:"column:level"`                         // 等级
	NumberFans       int    `json:"number_fans" gorm:"column:number_fans"`             // 粉丝数量
	NumberFollow     int    `json:"number_follow" gorm:"column:number_follow"`         // 关注博主的数量
}
