package vo

type Video struct {
	Vid              int     `json:"vid"`               // 视频id
	Uid              int     `json:"uid" `              // 视频作者id
	Examine          int     `json:"examine"`           // 视频审批
	VUrl             string  `json:"v_url"`             // 视频地址
	CidOne           int     `json:"cid_one"`           // 视频一级分类id
	CidTwo           int     `json:"cid_two"`           // 视频二级分类id
	Vimg             string  `json:"vimg"`              // 视频封面
	DurationTime     float32 `json:"duration_time"`     // 视频时长
	NumberLike       int     `json:"number_like"`       // 点赞量
	NumberCollect    int     `json:"number_collect" `   // 收藏量
	NumberCoin       int     `json:"number_coin" `      // 投币量
	Title            string  `json:"title"`             // 标题
	NumberViews      int     `json:"number_views"`      // 播放量
	BulletScreen     int     `json:"bullet_screen"`     // 弹幕量
	IntroduceBriefly string  `json:"introduce_briefly"` // 视频简介
	IsDelete         bool    `json:"is_delete"`         // 是否删除
	CreateTime       string  `json:"create_time"`       // 视频发布时间
	DeleteTime       string  `json:"delete_time"`       // 删除时间
}
