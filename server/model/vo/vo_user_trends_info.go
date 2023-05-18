package vo

type UserTrendsInfo struct {
	PushTime string `json:"push_time"` // 推送时间
	Uid      int    `json:"uid"`       // 用户id
	Vid      int    `json:"vid"`       // 视频id
}
