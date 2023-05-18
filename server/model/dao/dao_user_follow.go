package dao

type UserFollow struct {
	Uid        int    `json:"uid" gorm:"column:uid"`                 // 用户uid
	BUid       int    `json:"b_uid" gorm:"column:b_uid"`             // 博主uid
	FollowTime string `json:"follow_time" gorm:"column:follow_time"` // 关注时间
}

func (m *UserFollow) TableName() string {
	return "user_follow"
}
