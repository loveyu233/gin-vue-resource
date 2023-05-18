package dao

type UserTrendsInfo struct {
	PushTime string `json:"push_time" gorm:"column:push_time"` // 推送时间
	Uid      int    `json:"uid" gorm:"column:uid"`             // 用户id
	Vid      int    `json:"vid" gorm:"column:vid"`             // 视频id
}

func (m *UserTrendsInfo) TableName() string {
	return "user_trends_info"
}
