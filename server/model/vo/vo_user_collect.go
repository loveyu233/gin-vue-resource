package vo

type UserCollect struct {
	Uid        int    `json:"uid"`         // 用户id
	CollectVid int    `json:"collect_vid"` // 收藏的视频vid
	CreateTime string `json:"create_time"` // 收藏视频时间
}
