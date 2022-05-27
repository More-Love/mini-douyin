package relation

type RelationListResponse struct {
	StatusCode string  `json:"status_code"`// 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"` // 返回状态描述
	UserList   []User  `json:"user_list"`  // 用户信息列表
}

// User
type User struct {
	FollowCount   int64  `json:"follow_count"`  // 关注总数
	FollowerCount int64  `json:"follower_count"`// 粉丝总数
	ID            int64  `json:"id"`            // 用户id
	IsFollow      bool   `json:"is_follow"`     // true-已关注，false-未关注
	Name          string `json:"name"`          // 用户名称
}