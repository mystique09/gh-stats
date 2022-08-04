package api

import (
	"net/http"

	"github.com/mystique09/gh-profile/ent"
)

type GithunUserInfo struct {
	Id       uint32 `json:"id"`
	Username string `json:"login"`
	NodeId   string `json:"node_id"`
}

type HttpResponse struct {
	response *http.Response
	err      error
}

type Server struct {
	db *ent.Client
}
