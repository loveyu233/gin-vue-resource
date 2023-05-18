package dto

type UserCollect struct {
	Uid        int    `json:"uid" gorm:"column:uid"`                 // 用户id
	CollectVid int    `json:"collect_vid" gorm:"column:collect_vid"` // 收藏的视频vid
	CreateTime string `json:"create_time" gorm:"column:create_time"` // 收藏视频时间
}
