package dto

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"` // 邮箱
	Password string `json:"password"  binding:"required"`   // 密码
	X        string `json:"x"`
}

type UserVerificationCode struct {
	Email string `json:"email" binding:"required,email"` // 邮箱
	Code  string `json:"code" binding:"required"`        //验证码
}

type UserEmail struct {
	Email string `json:"email"` // 邮箱
}

type UserVo struct {
	Uid                int    `json:"uid" `                // id
	Username           string `json:"username" `           // 名称
	BackgroundImg      string `json:"background_img" `     // 背景图
	SculptureHead      string `json:"sculpture_head"`      // 头像
	IntroduceBriefly   string `json:"introduce_briefly"`   // 个人简介
	NumberCoin         int    `json:"number_coin"`         // 硬币数量
	NumberBCoin        int    `json:"number_b_coin"`       // B币数量
	Level              int    `json:"level"`               // 等级
	Experience         int    `json:"experience"`          // 经验值
	ExperienceRequired int    `json:"experience_required"` //升到下一级所需经验值
	NumberFans         int    `json:"number_fans"`         // 粉丝数量
	NumberFollow       int    `json:"number_follow"`       // 关注博主的数量
}
