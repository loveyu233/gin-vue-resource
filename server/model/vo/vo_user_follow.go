package vo

type UserFollow struct {
	Uid        int    `json:"uid"`         // 用户uid
	BUid       int    `json:"b_uid"`       // 博主uid
	FollowTime string `json:"follow_time"` // 关注时间
}
