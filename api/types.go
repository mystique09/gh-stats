package api

type GithunUserInfo struct {
	Id       uint32 `json:"id"`
	Username string `json:"login"`
	NodeId   string `json:"node_id"`
}
