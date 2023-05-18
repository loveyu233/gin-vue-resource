package dto

type UserTrends struct {
	Uid        int    `json:"uid" gorm:"column:uid"`                 // 用户id
	TrendsTime string `json:"trends_time" gorm:"column:trends_time"` // 动态查看时间
}
