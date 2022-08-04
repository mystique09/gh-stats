package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/mystique09/gh-profile/ent/user"
	"github.com/narqo/go-badge"
)

func renderErrorBadge(c echo.Context, label string) error {
	badge, err := badge.RenderBytes(label, "0", "red")
	if err != nil {
		return c.String(400, err.Error())
	}
	return c.Blob(404, "image/svg+xml", badge)
}

func (s *Server) profileVisits(c echo.Context) error {
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
		renderErrorBadge(c, "User not found.")
		return nil
	}
	respData, err := ioutil.ReadAll(recv.response.Body)
	if err != nil {
		return c.String(400, err.Error())
	}

	var userInfo GithunUserInfo
	json.Unmarshal(respData, &userInfo)

	update_user, err := s.db.User.Update().Where(user.ID(userInfo.Id), user.NodeID(userInfo.NodeId), user.Username(userInfo.Username)).AddVisits(1).Save(c.Request().Context())
	if err != nil || update_user == 0 {
		s.db.User.Create().SetID(userInfo.Id).SetNodeID(userInfo.NodeId).SetUsername(userInfo.Username).SetVisits(0).Save(c.Request().Context())
		badge, _ := badge.RenderBytes(label, "0", badge.Color(color))
		return c.Blob(200, "image/svg+xml", badge)
	}

	get_user, err := s.db.User.Get(c.Request().Context(), userInfo.Id)
	if err != nil {
		return c.String(400, err.Error())
	}
	badge, err := badge.RenderBytes(label, strconv.Itoa(int(get_user.Visits)), badge.Color(color))
	if err != nil {
		return c.String(400, err.Error())
	}
	c.Response().Header().Set("Cache-Control", "public, max-age=0, no-cache, no-store, must-revalidate")
	return c.Blob(200, "image/svg+xml", badge)
}
