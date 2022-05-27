package relation

type FolloerListResponse struct {
	StatusCode string  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	UserList   []User  `json:"user_list"`   // 用户列表
}
