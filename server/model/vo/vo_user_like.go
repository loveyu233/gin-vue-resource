package vo

type UserLike struct {
	Uid        int    `json:"uid"`         // 用户id
	LikeVid    int    `json:"like_vid"`    // 喜欢的视频vid
	CreateTime string `json:"create_time"` // 点赞视频时间
}
