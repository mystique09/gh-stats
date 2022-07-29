package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/narqo/go-badge"
	"io/ioutil"
	"net/http"
)

func profileVisits(c echo.Context) error {
	name := c.Param("name")
	color := c.QueryParam("color")
	label := c.QueryParam("label")

	if color == "" {
		color = badge.ColorBlue.String()
	}

	if label == "" {
		label = "Profile Visits"
	}

	resp := make(chan *HttpResponse)

	go func() {
		r, e := http.Get("https://api.github.com/users/" + name)
		resp <- &HttpResponse{response: r, err: e}
	}()

	recv := <-resp
	if recv.err != nil || recv.response.StatusCode == 404 {
		badge, err := badge.RenderBytes("User not found", "0", "red")
		if err != nil {
			return c.String(400, err.Error())
		}
		return c.Blob(404, "image/svg+xml", badge)
	}
	respData, err := ioutil.ReadAll(recv.response.Body)
	if err != nil {
		return c.String(400, err.Error())
	}

	var userInfo GithunUserInfo
	json.Unmarshal(respData, &userInfo)

	badge, err := badge.RenderBytes(label, "0", badge.Color(color))
	if err != nil {
		return c.String(400, err.Error())
	}
	return c.Blob(200, "image/svg+xml", badge)
}
