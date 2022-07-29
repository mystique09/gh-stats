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

	if color == "" {
		color = badge.ColorBlue.String()
	}

	var resp *http.Response

	resp, err := http.Get("https://api.github.com/users/" + name)
	if err != nil || resp.StatusCode == 404 {
		badge, err := badge.RenderBytes("User not found", "0", "red")
		if err != nil {
			return c.String(400, err.Error())
		}
		return c.Blob(404, "image/svg+xml", badge)
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.String(400, err.Error())
	}

	var userInfo GithunUserInfo
	json.Unmarshal(respData, &userInfo)

	badge, err := badge.RenderBytes("profile-visits", "0", badge.Color(color))
	if err != nil {
		return c.String(400, err.Error())
	}
	return c.Blob(200, "image/svg+xml", badge)
}
