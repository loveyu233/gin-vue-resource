package dto

type VideoInfo struct {
	Vid              int     `json:"vid" gorm:"column:vid"`                             // 视频id
	Vname            string  `json:"vname" gorm:"column:vname"`                         // 视频标题
	VUrl             string  `json:"v_url" gorm:"column:v_url"`                         // 视频地址
	DurationTime     float32 `json:"duration_time" gorm:"column:duration_time"`         // 视频时长
	IntroduceBriefly string  `json:"introduce_briefly" gorm:"column:introduce_briefly"` // 视频简介
}
