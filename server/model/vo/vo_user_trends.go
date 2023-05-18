package vo

type UserTrends struct {
	Uid        int    `json:"uid"`          // 用户id
	TrendsTime string `json:"trends_time" ` // 动态查看时间
}
