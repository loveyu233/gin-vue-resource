package vo

type VideoInfo struct {
	Vid              int     `json:"vid"`               // 视频id
	Vname            string  `json:"vname"`             // 视频标题
	VUrl             string  `json:"v_url"`             // 视频地址
	DurationTime     float32 `json:"duration_time"`     // 视频时长
	IntroduceBriefly string  `json:"introduce_briefly"` // 视频简介
	NumberViews      int     `json:"number_views"`      // 视频播放量
}
