package user

type LoginResponse struct {
	StatusCode int64   `json:"status_code"`// 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"` // 返回状态描述
	Token      *string `json:"token"`      // 用户鉴权token
	UserID     *int64  `json:"user_id"`    // 用户id
}