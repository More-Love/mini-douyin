package publish

type PublishActionResponse struct {
	StatusCode int64   `json:"status_code"`
	StatusMsg  *string `json:"status_msg"` 
}