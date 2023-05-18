package dao

type Video struct {
	Vid              int     `json:"vid" gorm:"column:vid"`                             // 视频id
	Uid              int     `json:"uid" gorm:"column:uid"`                             // 视频作者id
	Examine          int     `json:"examine" gorm:"column:examine"`                     // 视频审批
	VUrl             string  `json:"v_url" gorm:"column:v_url"`                         // 视频地址
	CidOne           int     `json:"cid_one" gorm:"column:cid_one"`                     // 视频一级分类id
	CidTwo           int     `json:"cid_two" gorm:"column:cid_two"`                     // 视频二级分类id
	Vimg             string  `json:"vimg" gorm:"column:vimg"`                           // 视频封面
	DurationTime     float32 `json:"duration_time" gorm:"column:duration_time"`         // 视频时长
	NumberLike       int     `json:"number_like" gorm:"column:number_like"`             // 点赞量
	NumberCollect    int     `json:"number_collect" gorm:"column:number_collect"`       // 收藏量
	NumberCoin       int     `json:"number_coin" gorm:"column:number_coin"`             // 投币量
	Title            string  `json:"title" gorm:"column:title"`                         // 标题
	NumberViews      int     `json:"number_views" gorm:"column:number_views"`           // 播放量
	BulletScreen     int     `json:"bullet_screen" gorm:"column:bullet_screen"`         // 弹幕量
	IntroduceBriefly string  `json:"introduce_briefly" gorm:"column:introduce_briefly"` // 视频简介
	IsDelete         bool    `json:"is_delete" gorm:"column:is_delete"`                 // 是否删除
	CreateTime       string  `json:"create_time" gorm:"column:create_time"`             // 视频发布时间
	DeleteTime       string  `json:"delete_time" gorm:"column:delete_time"`             // 删除时间
}

func (m *Video) TableName() string {
	return "video"
}
