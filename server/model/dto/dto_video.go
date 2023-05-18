package dto

type Video struct {
	Vid              int    `json:"vid" gorm:"column:vid"`                             // 视频id
	Uid              int    `json:"uid" gorm:"column:uid"`                             // 视频作者id
	CidOne           int    `json:"cid_one" gorm:"column:cid_one"`                     // 视频一级分类id
	CidTwo           int    `json:"cid_two" gorm:"column:cid_two"`                     // 视频二级分类id
	Vimg             string `json:"vimg" gorm:"column:vimg"`                           // 视频封面
	Title            string `json:"title" gorm:"column:title"`                         // 标题
	IntroduceBriefly string `json:"introduce_briefly" gorm:"column:introduce_briefly"` // 视频简介
	CreateTime       string `json:"create_time" gorm:"column:create_time"`             // 视频发布时间
}
