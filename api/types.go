package api

import "net/http"

type GithunUserInfo struct {
	Id       uint32 `json:"id"`
	Username string `json:"login"`
	NodeId   string `json:"node_id"`
}

type HttpResponse struct {
	response *http.Response
	err      error
}
