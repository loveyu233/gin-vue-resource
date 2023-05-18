package dao

type VideoClassify struct {
	Cid          int    `json:"cid" gorm:"column:cid"`                       // 分类id
	OneID        int    `json:"one_id" gorm:"column:one_id"`                 // 一级分类id
	ClassOneName string `json:"class_one_name" gorm:"column:class_one_name"` // 视频一级分类名称
	TwoID        int    `json:"two_id" gorm:"column:two_id"`                 // 二级分类id
	ClassTwoName string `json:"class_two_name" gorm:"column:class_two_name"` // 视频二级分类名称
}

func (m *VideoClassify) TableName() string {
	return "video_classify"
}
